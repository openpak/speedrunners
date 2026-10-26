// Command speedrunnersd is OpenPak's SpeedRunners (Switch) title front door.
//
// SpeedRunners (0100C8B00C4AE000, Unity 2018.3.10f1 IL2CPP) is a pure NEX
// title: NEX 4.6.3 (DataStore, MatchMaking, Ranking, Ranking2, Subscriber,
// Utility plus Authentication, SecureConnection, NATTraversal) over
// PRUDP/PRUDPLite, with a PiaNex 5.18.0 session+lan path for local wireless.
// It reaches no developer backend — matchmaking discovery uses the standard
// NEX hosts n/p/s.n.srv.nintendo.net — so there is no game HTTP API to serve
// here. Online play is served by platforms/switch/games/switch-nex once this
// title is registered there; see docs/nex-registration.md.
//
// This daemon holds the title's runtime configuration for that stack: which
// switch-nex auth/secure endpoints to point the console at, and the game's
// NEX access key once it is recovered (TODO-KEY: only the `accessKey`
// parameter name survives in the binary's strings; the value needs a Ghidra
// xref in main — see docs/protocol.md). Until the key lands the title cannot
// authenticate and /healthz says so plainly.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	cfg := configFromEnv()
	client, err := newNEXClient(cfg)
	if err != nil {
		// The daemon still serves: /healthz reports the missing key so a
		// readiness check can see the title is not playable yet.
		log.Printf("[speedrunners] NEX client not ready: %v", err)
	} else {
		log.Printf("[speedrunners] NEX client ready: auth=%s secure=%s", client.authURL(), client.secureURL())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, healthLine(cfg)+"\n")
	})
	log.Printf("[speedrunners] serving on %s", cfg.addr)
	log.Fatal(http.ListenAndServe(cfg.addr, mux))
}

// healthLine is the /healthz body. The key= field is the readiness signal:
// "missing" until the access key file holds the recovered value (TODO-KEY).
func healthLine(cfg config) string {
	key := "missing"
	if cfg.accessKey != "" {
		key = "present"
	}
	return fmt.Sprintf("ok speedrunners auth=%s secure=%s key=%s", cfg.nexAuth, cfg.nexSecure, key)
}

// config is the daemon's env-driven configuration.
type config struct {
	addr      string // HTTP listen address (serves /healthz)
	nexAuth   string // switch-nex auth endpoint for this title (host:port)
	nexSecure string // switch-nex secure endpoint for this title (host:port)
	accessKey string // NEX game access key value (TODO-KEY: outstanding)
}

func configFromEnv() config {
	cfg := config{
		addr:      envOr("SPEEDRUNNERS_ADDR", ":23610"),
		nexAuth:   envOr("SPEEDRUNNERS_NEX_AUTH", "127.0.0.1:22360"),
		nexSecure: envOr("SPEEDRUNNERS_NEX_SECURE", "127.0.0.1:22361"),
	}
	if path := os.Getenv("SPEEDRUNNERS_ACCESS_KEY_FILE"); path != "" {
		if raw, err := os.ReadFile(path); err != nil {
			log.Printf("[speedrunners] cannot read access key file %s: %v", path, err)
		} else {
			cfg.accessKey = strings.TrimSpace(string(raw))
		}
	} else if v := os.Getenv("SPEEDRUNNERS_ACCESS_KEY"); v != "" {
		cfg.accessKey = v
	}
	return cfg
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

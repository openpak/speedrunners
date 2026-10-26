package main

// Middleware client stub for SpeedRunners' NEX stack.
//
// The game speaks NEX 4.6.3 over PRUDP/PRUDPLite to the standard NEX
// discovery hosts; OpenPak answers that traffic from switch-nex once this
// title is registered (docs/nex-registration.md). This stub models the
// client's side of that connection — the station URLs the console dials and
// the access key it authenticates with — and refuses to build a client until
// the outstanding key material lands (TODO-KEY). It dials nothing yet: the
// first connection attempt belongs to the registration bring-up, not to this
// skeleton.

import (
	"errors"
	"fmt"
	"strings"
)

// TitleID is SpeedRunners on Switch (USA v0 dump).
const TitleID = "0100C8B00C4AE000"

// NEXVersion is the NEX release the game's main binary links.
const NEXVersion = "4.6.3"

// ErrAccessKeyMissing is returned while the game's NEX access key value is
// still outstanding (TODO-KEY).
var ErrAccessKeyMissing = errors.New("speedrunners: NEX access key outstanding (TODO-KEY)")

// nexClient is the stub NEX middleware client: endpoint configuration plus
// the key, ready for the PRUDP dialer the registration bring-up will add.
type nexClient struct {
	cfg config
}

// newNEXClient validates the title configuration. It succeeds only when the
// access key is present; the daemon stays up without one so /healthz can
// report the gap.
func newNEXClient(cfg config) (*nexClient, error) {
	if strings.TrimSpace(cfg.accessKey) == "" {
		return nil, ErrAccessKeyMissing
	}
	if cfg.nexAuth == "" || cfg.nexSecure == "" {
		return nil, fmt.Errorf("speedrunners: NEX auth/secure endpoints unset")
	}
	return &nexClient{cfg: cfg}, nil
}

// authURL and secureURL are the station-URL forms of the configured NEX
// endpoints — the addresses the console dials (raw IP:port, no DNS).
func (c *nexClient) authURL() string {
	return "prudps:/address=" + c.cfg.nexAuth + ";sid=1"
}

func (c *nexClient) secureURL() string {
	return "prudps:/address=" + c.cfg.nexSecure + ";sid=1"
}

package main

import (
	"errors"
	"os"
	"testing"
)

func TestNewNEXClientNeedsAccessKey(t *testing.T) {
	cfg := config{addr: ":23570", nexAuth: "127.0.0.1:22360", nexSecure: "127.0.0.1:22361"}
	if _, err := newNEXClient(cfg); !errors.Is(err, ErrAccessKeyMissing) {
		t.Fatalf("newNEXClient without key = %v, want ErrAccessKeyMissing", err)
	}
	cfg.accessKey = "test-key-from-file"
	c, err := newNEXClient(cfg)
	if err != nil {
		t.Fatalf("newNEXClient with key = %v", err)
	}
	if got := c.authURL(); got != "prudps:/address=127.0.0.1:22360;sid=1" {
		t.Fatalf("authURL = %q", got)
	}
	if got := c.secureURL(); got != "prudps:/address=127.0.0.1:22361;sid=1" {
		t.Fatalf("secureURL = %q", got)
	}
}

func TestHealthLineReportsKeyState(t *testing.T) {
	cfg := config{nexAuth: "127.0.0.1:22360", nexSecure: "127.0.0.1:22361"}
	if got := healthLine(cfg); got != "ok speedrunners auth=127.0.0.1:22360 secure=127.0.0.1:22361 key=missing" {
		t.Fatalf("healthLine without key = %q", got)
	}
	cfg.accessKey = "k"
	if got := healthLine(cfg); got != "ok speedrunners auth=127.0.0.1:22360 secure=127.0.0.1:22361 key=present" {
		t.Fatalf("healthLine with key = %q", got)
	}
}

func TestConfigFromEnvDefaults(t *testing.T) {
	t.Setenv("SPEEDRUNNERS_ADDR", "")
	t.Setenv("SPEEDRUNNERS_NEX_AUTH", "")
	t.Setenv("SPEEDRUNNERS_NEX_SECURE", "")
	t.Setenv("SPEEDRUNNERS_ACCESS_KEY", "")
	t.Setenv("SPEEDRUNNERS_ACCESS_KEY_FILE", "")
	cfg := configFromEnv()
	if cfg.addr != ":23570" || cfg.nexAuth != "127.0.0.1:22360" || cfg.nexSecure != "127.0.0.1:22361" {
		t.Fatalf("defaults = %+v", cfg)
	}

	t.Setenv("SPEEDRUNNERS_ACCESS_KEY", "  hunter2  ")
	// Value from env is used verbatim (files are trimmed, env is explicit).
	if cfg := configFromEnv(); cfg.accessKey != "  hunter2  " {
		t.Fatalf("env key = %q", cfg.accessKey)
	}

	f, err := os.CreateTemp(t.TempDir(), "key")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString("file-key\n"); err != nil {
		t.Fatal(err)
	}
	f.Close()
	t.Setenv("SPEEDRUNNERS_ACCESS_KEY", "")
	t.Setenv("SPEEDRUNNERS_ACCESS_KEY_FILE", f.Name())
	if cfg := configFromEnv(); cfg.accessKey != "file-key" {
		t.Fatalf("file key = %q", cfg.accessKey)
	}
}

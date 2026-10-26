# Changelog — speedrunners

## v0.1.0 — unreleased

- `speedrunnersd`: title front door for SpeedRunners (`0100C8B00C4AE000`, NEX
  4.6.3 + PiaNex 5.18.0). Env-driven config (NEX auth/secure endpoints,
  access key file), `/healthz` readiness with a `key=` signal, and a NEX
  middleware client stub that stays unbuilt until the access key lands
  (TODO-KEY).
- `docs/nex-registration.md`: staged switch-nex registration spec (title ID,
  NEX version, protocol set, requested port base 22360, outstanding keys).

# SpeedRunners

OpenPak's SpeedRunners Server (Nintendo Switch).

| | |
| --- | --- |
| Game | SpeedRunners |
| Network | Nintendo Switch |
| Title ID | `0100C8B00C4AE000` |
| Licence | AGPL-3.0-only |

## What it serves

SpeedRunners is a pure NEX title: online play is NEX 4.6.3 over
PRUDP/PRUDPLite plus a PiaNex 5.18.0 session+lan path for local wireless.
There is no developer backend — matchmaking uses the standard NEX discovery
hosts `n/p/s.n.srv.nintendo.net`.

| Host | What | Served by |
| --- | --- | --- |
| `n/p/s.n.srv.nintendo.net` | Auth, matchmaking, DataStore, Ranking/Ranking2, NAT traversal (NEX) | [`switch-nex`](https://github.com/openpak/switch-nex), once this title is registered |
| local wireless (`ldn:u`) | LAN multiplayer (Pia session+lan) | console to console; nothing to host |

This repo (`speedrunnersd`) is the title front door: it holds the runtime
configuration for that stack and reports readiness on `/healthz`. The title
is not playable until the NEX access key is recovered (TODO-KEY) and the
[registration spec](docs/nex-registration.md) is applied to switch-nex.

## Run

```sh
go run ./cmd/speedrunnersd
```

| Variable | Meaning |
| --- | --- |
| `SPEEDRUNNERS_ADDR` | HTTP listen address (default `:23610`; serves `/healthz`) |
| `SPEEDRUNNERS_NEX_AUTH` | switch-nex auth endpoint for this title (default `127.0.0.1:22360`) |
| `SPEEDRUNNERS_NEX_SECURE` | switch-nex secure endpoint for this title (default `127.0.0.1:22361`) |
| `SPEEDRUNNERS_ACCESS_KEY_FILE` | file holding the NEX access key value (outstanding — TODO-KEY) |
| `SPEEDRUNNERS_ACCESS_KEY` | the key value directly (overridden by the file when both are set) |

`/healthz` answers `ok speedrunners auth=… secure=… key=present|missing`; the
`key=` field is the readiness signal.

## Docs

- [Protocol](docs/protocol.md)
- [NEX registration spec](docs/nex-registration.md)
- [Ports](docs/ports.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).

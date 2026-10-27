# NEX registration spec — SpeedRunners

Staged here for the serial switch-nex integration; this repo never edits
`platforms/switch/games/switch-nex/`. The registrar copies this entry into
the per-title registry (`cmd/switch-nexd`) and `titles/` once the
outstanding key material lands.

Wire facts sourced from `scratch/decompile-speedrunners/RECON.md` and
`scratch/decompile-speedrunners/findings.json`; see
[protocol.md](protocol.md).

| Field | Value |
| --- | --- |
| Title | SpeedRunners |
| Title ID | `0100C8B00C4AE000` |
| NEX version | 4.6.3 (`NEX_DS/MM/R2/RK/SS/UT-4_6_3`) |
| Pia | PiaNex 5.18.0 (session+lan) |
| Transport | PRUDP + PRUDPLite (`prudp:/`, `prudps:/`) |
| Protocols | Auth, SecureConnection, TicketGranting, NATTraversal (+Relay, ReportInternal), MatchMaking (+Ext, MatchmakeExtension), DataStore (+LogicServer), Ranking, Ranking2, Subscriber, Utility, Notification/NintendoNotificationEvent servers, Dns, Http, FileServer, AcquireNsaIdToken |
| Requested port base | **22360** (auth 22360, secure 22361, healthz 22362) — free; neighbours 22350 (SSBU) and 22370 (3D World) are taken |
| Game server ID | `2C186B00` (recovered 2026-10-27, see below) |
| Access key | `ae8528ab` (recovered 2026-10-27, see below) |
| Outstanding | PRUDP game port and gathering/DataStore usage profile pending RMC capture on first routed boot |
| Registry status | **REGISTERED** 2026-10-27: `switch-nexd` entry `{"speedrunners", 22360, "ae8528ab", ...}` + `bin/title-ports.json`; provenance in `switch-nex/docs/speedrunners.md` |

## Key recovery (2026-10-27)

Il2CppDumper on the sweep `main.nso` + `global-metadata.dat`
(metadata v24.1) exposes named managed constants in
`Assets.Switch.SwitchOnline`: `GAME_SERVER_ID = 739797760`
(`0x2C186B00`), `GAME_SERVER_KEY = "ae8528ab"`. Native
`InitFrameworkTask::Execute` (`0x34d8a0`) builds the ID into `w23`
(`0x34db7c`/`0x34db80`, stored `[x21+0x1c]`) and the key string
object (`0x3081E38`) at `[x21+0x20]` — same-call construction.
Harbor reaches the online account gate; wire confirmation awaits a
linked test account.

## Registration notes

- Cleanest registration path of the batch: pure NEX title, standard
  discovery hosts (`n/p/s.n.srv.nintendo.net`), no third-party backend to
  stub, no custom hosts to redirect.
- The LAN path (`nn::pia::lan` + `ldn:u` + `Lan*ProcessHostMigrationJob`)
  is console-to-console; nothing to register for it.
- `Lidgren.Network.Unity.dll` is metadata-only (no DLL ships); ignore unless
  a capture proves otherwise.

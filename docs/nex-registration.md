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
| Outstanding keys | **TODO-KEY**: NEX game access key VALUE (parameter name only in strings; needs Ghidra xref in `main`). PRUDP game port and gathering/DataStore usage profile likewise pending IL2CPP dump + RMC capture |

## Registration notes

- Cleanest registration path of the batch: pure NEX title, standard
  discovery hosts (`n/p/s.n.srv.nintendo.net`), no third-party backend to
  stub, no custom hosts to redirect.
- The LAN path (`nn::pia::lan` + `ldn:u` + `Lan*ProcessHostMigrationJob`)
  is console-to-console; nothing to register for it.
- `Lidgren.Network.Unity.dll` is metadata-only (no DLL ships); ignore unless
  a capture proves otherwise.

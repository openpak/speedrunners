# SpeedRunners (Switch) — wire facts

Every fact below comes from the workspace recon of the Switch dump —
`scratch/decompile-speedrunners/RECON.md` and
`scratch/decompile-speedrunners/findings.json` (USA v0 NSP,
`SpeedRunners[0100C8B00C4AE000][USA][v0].nsp`). Nothing here is guessed from
other titles. Keep this file in step with `cmd/speedrunnersd`.

## Build

- Title ID `0100C8B00C4AE000`, SDK 7.4.0.0 (Meta built with 9.3.0.0).
- Engine: Unity 2018.3.10f1, IL2CPP (`global-metadata.dat`,
  `Managed/Metadata`), `scripting-runtime-version=legacy`.
- ExeFS: `main` (25 MB NSO: game + NEX + Unity native), `sdk` (5.3 MB
  Nintendo SDK), `subsdk0` (3.5 MB Nintendo multimedia — video, not
  network), `rtld`.
- RomFS is `Data/` only (575 MB game assets; no server configs, no hosts).
- NPDM network services: `bsd:u`, `nsd:u`, `nifm:u`, `sfdnsres`, `ssl`,
  `ldn:u` (local wireless), `friend:u`, `bcat:u`, `pctl`, `set`, `time:u`.

## Network middleware: NEX 4.6.3 (+ PIA), no third-party backend

Version strings in `main` (`RECON.md` §"Network middleware"):

- `SDK MW+Nintendo+NEX-4_6_3-`, `NEX_DS/MM/R2/RK/SS/UT-4_6_3`,
  `PiaNex-5_18_0-forNEX-4_6_2`.

NEX protocol clients linked, all in `main`:

- Authentication, SecureConnection, TicketGranting, NATTraversal (+ Relay,
  ReportInternal), MatchMaking (+ Ext, MatchmakeExtension), DataStore (+
  LogicServer), Ranking + Ranking2, Subscriber, Utility, MessageDelivery-side
  Notification / NintendoNotificationEvent servers, Dns, Http, FileServer,
  AcquireNsaIdToken.

Transport and hosts:

- PRUDP + PRUDPLite (`nn::nex::PRUDPEndPoint/Stream`, `prudp:/`,
  `prudps:/`).
- Matchmaking discovery uses `n/p/s.n.srv.nintendo.net` — the standard NEX
  hosts. No custom game hosts, no hardcoded game-server IPs (loopback and
  generics only).

Local wireless path:

- `nn::pia::session`, `nn::pia::lan` plus the `ldn:u` service.

Game's own netcode layer:

- `OnlineCore` (build paths `.../OnlineCore/src/...`,
  `E:/Nintendo/Version7_3_2/SpeedRunners/...`) with `Nex*` /
  `Lan*ProcessHostMigrationJob` host-migration jobs for both NEX and LAN.
- Managed (C#) wrapper: `AcquireNexUniqueId*`,
  `AssociateNexUniqueIdWithMyPrincipalId*`, lobby API (`CreateAndJoinLobby`,
  `LobbyEntered/Refused`, `BlackListLobby`), `RankingClient::*`,
  `CSpeedRunner.PlatformSwitch`.

## What is not here

- No Photon, PlayFab, EOS, Demonware, or Steamworks (0 hits in
  `findings.json` `middleware.absent`).
- `Lidgren.Network.Unity.dll` appears as an assembly name in metadata but no
  `.dll` ships in `Managed/` — a dead PC/XNA leftover, presumed unused.
- TLS via the `ssl` service + libcurl (cookie docs URL only); no bundled
  cert strings.

## Outstanding (TODO-KEY and follow-ups)

From `RECON.md` §"Gaps / next steps" and `findings.json` `unresolved`:

1. **NEX game access key value** — only the `accessKey` parameter name
   survives in strings; needs a Ghidra xref in `main` (decompressed copy at
   `scratch/decompile-speedrunners/nso/main.uncomp`). Until it lands the
   title cannot authenticate; `speedrunnersd` reports `key=missing`.
2. **PRUDP game port + gathering/DataStore usage profile** — which
   DataStore/Ranking2 calls and gathering types the game uses.
3. **IL2CPP signature dump** (`global-metadata.dat` + `main`) for the C#
   signatures, then an RMC capture plan.
4. **No switch-nex registry entry yet** — registration needs the key
   material from (1); the spec is staged in
   [nex-registration.md](nex-registration.md).

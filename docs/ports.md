# Ports

| Port | Use |
| --- | --- |
| 23610 | plain HTTP: `/healthz` (this repo, `speedrunnersd`) |
| 23611–23619 | reserved |
| 22360–22362 | requested NEX block on switch-nex (auth 22360, secure 22361, healthz 22362); see [nex-registration.md](nex-registration.md) |

Online play uses the NEX block only; this service publishes no game port.

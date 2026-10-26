# syntax=docker/dockerfile:1
# SpeedRunners (Switch) title front door. Plain HTTP on 23610 (serves /healthz);
# online play itself is NEX on switch-nex (requested port base 22360).
FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build
ARG TARGETOS TARGETARCH
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOWORK=off go build -trimpath -buildvcs=false -ldflags="-s -w" -o /out/speedrunnersd ./cmd/speedrunnersd

FROM gcr.io/distroless/static
COPY --from=build /out/speedrunnersd /speedrunnersd
USER nonroot:nonroot
EXPOSE 23610
ENTRYPOINT ["/speedrunnersd"]

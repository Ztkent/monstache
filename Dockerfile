FROM --platform=$BUILDPLATFORM golang:1.22.10-alpine3.21 AS build
WORKDIR /src
ARG TARGETOS TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
	go mod download; \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/monstache .

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/bin/monstache"]
COPY --from=build /out/monstache /bin

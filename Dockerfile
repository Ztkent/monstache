# Build the app and plugin
FROM golang:1.21 AS build-plugin

RUN mkdir /app
WORKDIR /app

COPY ./monstache ./
COPY ./monstache_config.toml ./
COPY ./monstache/plugin ./

RUN make release
RUN go mod download
RUN go build -buildmode=plugin -o /app/build/plugin.so plugin.go

# Copy the files we need and run the app
FROM golang:1.21 AS monstache

COPY --from=build-plugin /app/build/linux-amd64/monstache /bin/monstache
COPY --from=build-plugin /app/build/plugin.so /bin/plugin.so
COPY --from=build-plugin /app/monstache_config.toml /etc/monstache/monstache.toml

ENTRYPOINT /bin/monstache -f /etc/monstache/monstache.toml -mongo-url ${MONGODB_URI} -elasticsearch-url ${OPENSEARCH_HOST} -elasticsearch-user ${OPENSEARCH_USER} -elasticsearch-password ${OPENSEARCH_PASS}


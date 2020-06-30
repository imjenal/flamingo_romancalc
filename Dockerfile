FROM golang:1.12 AS builder
WORKDIR /deeptrace
COPY . /deeptrace/

WORKDIR /deeptrace/
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -ldflags "-w -s -X main.Version=${VERSION} -X main.MinVersion=$(git rev-parse HEAD) -X main.BuildTime=$(date +%FT%T%z)" ./app/server.go

# copying builder to final
FROM alpine:3.10.2 as deploy
WORKDIR /deeptrace/

COPY --from=builder /deeptrace/server ./server

EXPOSE 8081
CMD [ "./server"]
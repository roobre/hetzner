FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /hetzner ./cmd

FROM alpine:latest

RUN apk --no-cache add tini
COPY --from=builder /hetzner /hetzner
ENTRYPOINT ["/sbin/tini", "--", "/hetzner"]

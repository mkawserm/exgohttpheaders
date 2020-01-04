FROM golang:1.13.5 as builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app main.go


FROM alpine:latest

# mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app

WORKDIR /app
EXPOSE 8080

COPY --from=builder /app/app .

ENTRYPOINT ["/app/app"]

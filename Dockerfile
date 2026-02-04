## Builder Stage
FROM --platform=linux/arm64 golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
    -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -trimpath \
    -o ./dist/api ./cmd/main.go

## Api Stage
FROM --platform=linux/arm64 alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/dist/api .
COPY --from=builder /app/internal/handlers/spec ./internal/handlers/spec
EXPOSE 8080
ENTRYPOINT ["./api"]
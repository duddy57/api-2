FROM golang:1.25-alpine as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags='-w -s -extldflags "-static"' -trimpath -o /app/bin/api ./cmd

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/bin/api .
COPY --from=builder /app/internal/handlers/spec ./internal/handlers/spec
EXPOSE 8080
ENTRYPOINT ["./api"]
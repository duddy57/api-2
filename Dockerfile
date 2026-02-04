## Builder Stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build \
    -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -trimpath \
    -o ./bin/api ./cmd/main.go

## Api Stage
FROM scratch
WORKDIR /app
COPY --from=builder ./bin/api .
EXPOSE 8080
ENTRYPOINT ["./api"]
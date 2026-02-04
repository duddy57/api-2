FROM golang:1.24.0-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /bin/api ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /bin/api .
COPY --from=builder /app/internal/handlers/spec ./internal/handlers/spec
EXPOSE 8080
ENTRYPOINT ["./api"]
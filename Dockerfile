FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/app ./cmd/app

FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY --from=builder /app/bin/app /app/app

EXPOSE 8080

USER nonroot:nonroot
CMD ["/app/app"]
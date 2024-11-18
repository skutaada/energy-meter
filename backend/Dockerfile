FROM golang:1.23.3 AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /backend .

FROM gcr.io/distroless/base-debian12:latest

COPY --from=base /backend .

ENTRYPOINT ["./backend"]

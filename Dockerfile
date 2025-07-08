FROM golang:alpine AS build

WORKDIR /buildapp

COPY . .

RUN go build -o gocrud main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build /buildapp/gocrud /app/gocrud
COPY .env /app/.env

ENTRYPOINT ["/app/gocrud"]

FROM golang:1.23-bullseye

WORKDIR /wsserver

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", "ws.air.toml"]

FROM golang:1.23-bullseye

WORKDIR /wsserver

RUN go install github.com/air-verse/air@latest

CMD ["sh", "-c", "cd ./server && air -c ws.air.toml"]

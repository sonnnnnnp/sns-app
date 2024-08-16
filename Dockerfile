FROM golang:1.22-bullseye

WORKDIR /api

RUN go install github.com/air-verse/air@latest
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

CMD ["air"]

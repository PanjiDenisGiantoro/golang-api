FROM golang:alpine

RUN apk update && apk add--no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

run go build -o binary

ENTRYPOINT ["app/binary"]
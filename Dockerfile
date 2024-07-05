FROM golang:1.20-alpine

WORKDIR /usr/src/app

COPY src .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

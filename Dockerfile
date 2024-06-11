FROM golang:1.22.3-alpine

WORKDIR /usr/src/app

COPY src .
RUN go build -v -o /usr/local/bin/app ./...

# EXPOSE 9001
CMD ["app"]

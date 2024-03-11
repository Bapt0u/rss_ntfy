FROM golang:1.22.1 as builder

WORKDIR /src
COPY . /src

RUN go build -o rssfeeder main.go

FROM alpine:latest

COPY --from=builder /src/rssfeeder /bin/rssfeeder
CMD ["/bin/rssfeeder"]
FROM golang:1.16 as builder

COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o /build/main .

FROM alpine:3.10

COPY --from=builder /build/main /bin/main
RUN chmod +x /bin/main

CMD ["/bin/main"]

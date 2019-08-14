FROM golang:1.12.7 as builder

# Add Maintainer Info
LABEL maintainer="Magne Skjeret <magne.skjeret@gmail.com>"

WORKDIR /go/src/github.com/mskjeret/ingress-http-poller
COPY . .
RUN go get -u github.com/rancher/trash && trash
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ingress-http-poller .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/mskjeret/ingress-http-poller/ingress-http-poller .

CMD ["./ingress-http-poller"]
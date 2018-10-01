FROM golang:alpine as builder
WORKDIR /go-expvar
RUN apk add --no-cache git
COPY . /go-expvar
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:edge
COPY --from=builder /go-expvar/go-expvar /go-expvar
EXPOSE 8080
CMD ["/go-expvar"]

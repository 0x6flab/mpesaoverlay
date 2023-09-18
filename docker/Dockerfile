FROM golang:1.20-alpine AS builder
LABEL org.opencontainers.image.source https://github.com/0x6flab/mpesaoverlay
ARG SVC
ARG GOARCH
ARG GOARM
ARG GOOS
ARG CGO_ENABLED
ARG VERSION
ARG COMMIT
ARG TIME

WORKDIR /go/src/github.com/0x6flab/mpesaoverlay
COPY . .
RUN apk update \
    && apk add make\
    && make $SVC \
    && mv build/mpesa-$SVC /exe

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]

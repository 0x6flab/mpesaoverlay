FROM golang:1.22-alpine AS builder
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
    && go mod vendor \
    && make $SVC \
    && mv build/mpesa-$SVC /exe

FROM scratch
LABEL org.opencontainers.image.source https://github.com/0x6flab/mpesaoverlay
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]

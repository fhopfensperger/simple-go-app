FROM golang:1.24 as goBuilder

USER root
WORKDIR /work
COPY . .
ARG BUILD_VERSION="0.0.0"
RUN CGO_ENABLED=0 go build -a -ldflags "-X main.version=$BUILD_VERSION" -o simple-go-app .

FROM alpine:3.21

LABEL maintainer="Florian Hopfensperger <f.hopfensperger@gmail.com>"

RUN apk add --update wget curl git openssl ca-certificates \
    && rm /var/cache/apk/* \
    && adduser -G root -u 1000 -D -S kuser

USER 1000
WORKDIR /app

COPY --chown=1000:0 --from=goBuilder /work/simple-go-app .

EXPOSE 8080
ENTRYPOINT ["./simple-go-app"]

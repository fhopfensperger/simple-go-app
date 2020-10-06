FROM golang:alpine AS build
ADD . /build
WORKDIR /build
ARG BUILD_VERSION="v0.0.0"
RUN go build -ldflags "-X main.version=$BUILD_VERSION" -o app ./...

# 2. stage
FROM alpine

RUN apk --no-cache add --update curl \
    && apk update \
    && apk upgrade 

COPY --from=build /build/app /usr/bin/app

EXPOSE 8080
CMD [ "/usr/bin/app" ]
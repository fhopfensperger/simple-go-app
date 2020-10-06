FROM golang:alpine AS build
ADD . /build
WORKDIR /build
RUN go build -o app ./...
# RUN go build -o app ./... -ldflags "-X main.version=v2.8.123"

# 2. stage
FROM alpine

RUN apk --no-cache add --update curl \
    && apk update \
    && apk upgrade 

COPY --from=build /build/app /usr/bin/app

EXPOSE 8080
CMD [ "/usr/bin/app" ]
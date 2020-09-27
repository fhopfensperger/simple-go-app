FROM golang:alpine AS build
ADD . /build
WORKDIR /build
RUN go build -o app ./...

# 2. stage
FROM alpine
COPY --from=build /build/app /usr/bin/app

EXPOSE 8080
CMD [ "/usr/bin/app" ]
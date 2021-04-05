# simple-go-app

A really simple rest web service written in Go

## Building
```bash
docker buildx build \
  --platform linux/arm/v7,linux/amd64,linux/arm64 \
  --tag quay.io/fhopfensperger/simple-go-app \
  --tag quay.io/fhopfensperger/simple-go-app:v0.1.0 \
  -f Dockerfile \
  --build-arg BUILD_VERSION=v0.1.0 \
  --push .
```
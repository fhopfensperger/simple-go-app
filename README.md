[![Docker Repository on Quay](https://quay.io/repository/fhopfensperger/simple-go-app/status "Docker Repository on Quay")](https://quay.io/repository/fhopfensperger/simple-go-app)

# simple-go-app

A really simple rest web service written in Go.

After starting it exposes the following `GET` endpoints:
- `/hello` -> returns a string response
- `/hello-json` -> returns a json response
- `/metrics` -> returns with Prometheus standard metrics
- `/health/live` -> 200 OK
- `/health/ready` -> 200 OK

The port ist set to `8080` and can be controlled with an environment variable `PORT`

Json logging can be enabled with the environment variable `JSON_LOG`

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

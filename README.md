# simple-go-app

A really simple rest web service written in Go.

After starting it exposes the following `GET` endpoints:
- `/hello` -> returns a string response
- `/hello-json` -> returns a json response
- `/metrics` -> returns Prometheus standard metrics
- `/health/live` -> 200 OK
- `/health/ready` -> 200 OK

The port ist set to `8080` and can be controlled with an environment variable `PORT`

Json logging can be enabled with the environment variable `JSON_LOG`


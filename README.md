```
     _                                 ____
    | | __ ___   ____ _  __   _____   / ___| ___
 _  | |/ _` \ \ / / _` | \ \ / / __| | |  _ / _ \
| |_| | (_| |\ V / (_| |  \ V /\__ \ | |_| | (_) |
 \___/ \__,_| \_/ \__,_|   \_/ |___/  \____|\___/

```

# Prerequisites

1. [Go](https://golang.org/) version >= 1.12
2. Plugin for Go in your favorite IDE (eg. [GoLand]( https://plugins.jetbrains.com/plugin/9568-go) for Intellij)
3. [ab](https://httpd.apache.org/docs/2.4/programs/ab.html) tool for load testing
4. [Docker](https://www.docker.com/) version >= 2.x

# Examples

## First Go application

1. Go to the project directory, and execute:

   ```bash
   go mod init com.jamf.services.java_vs_go
   ```
   
   It will initialize go project with `` which describes module name and version of go.
   It will be also used to manage project dependencies, like `pom.xml` or `build.gradle`.

1. Create directory and file for your application code

   ```bash
   mkdir cmd; cd cmd
   touch go_sample_service.go
   ```

1. Paste the following code in `go_sample_service.go`:

   ```go
   package main

   // public static void main(String[] args) {
   func main() {
     // argsWithProg := os.Args
     // argsWithoutProg := os.Args[1:]
     println("Hello Silesia Java Users!")
   }
   ```

## Simple HTTP service

Very simple HTTP microservice:
- the service listens on port 8080
- the service exposes one endpoint: `/hello` which returns the following string: `Hello Silesia Java Users!`

Checkout source code:

```bash
git checkout http_service
```

Run the service:

```bash
go run cmd/go_sample_service.go
```

## Echo HTTP service and testing

Echo service:
- the service exposes endpoint: `/echo-request`
- the endpoint returns response, which contains data from the request encoded in JSON format
- tests for the endpoint

Example of the response:

```json
{
  "queryString": "?test=true",
  "headers": [
    {
      "name": "content-type",
      "value": "application/json"
    },
    {
      "name": "accept",
      "value": "*/*"
    }
  ],
  "body": "body-as-string"
}
```

Checkout source code:

```bash
git checkout echo_service_and_testing
```

Run tests

```bash
go test ./... 
```

Run tests with coverage

```bash
go test ./... -cover
```

Run the service:

```bash
go run cmd/go_sample_service.go
```

## Load on HTTP service + memory usage

Generate some load and measure memory usage.

```bash
git checkout memory_usage
```

Instructions:

1. Run the service:

   ```bash
   go run cmd/go_sample_service.go
   ```

1. Open Activity Monitor and find process called `go_sample_service`.

1. Verify if application is working correctly:

   ```bash
   curl -H 'Content-Type: application/json' -X POST -d '@docs/assets/sample_payload.json' http://localhost:8080/echo-request
   ```

1. Generate some load on application:

   ```bash
   ab -T 'application/json' -p 'docs/assets/sample_payload.json' -n 10000 -c 100 http://localhost:8080/echo-request
   ```

## Simple HTTP client + stubbing

Proxy to `httpbin` service:
- the service exposes one endpoint: `/httpbin/json`
- the endpoint rewrites response returned by the following endpoint: http://httpbin.org/#/Response_formats/get_json.

Checkout source code:

```bash
git checkout http_client_and_stubbing
```

Run tests

```bash
go test ./... 
```

Run tests with coverage

```bash
go test ./... -cover
```

Run the service:

```bash
go run cmd/go_sample_service.go
```

## Dependency management with `gomod`

Refactoring of simple HTTP microservice:
- add [gorilla mux](https://github.com/gorilla/mux) dependency to the project and use it to simplify endpoints development

Checkout source code:

```bash
git checkout deps_management
```

Run the service:

```bash
go run cmd/go_sample_service.go
```

## Docker image

Build docker image with the service, and verify it's size.

Checkout source code:

```bash
git checkout docker_image
```

Instructions:

1. Build image for this service:

   ```bash
   docker build -t go_sample_service:1.0.0 -f build/package/Dockerfile_service_go ./
   ```

1. Verify the service image size:

   ```bash
   docker images | grep go_sample_service
   ```
   
1. Run the service image:

   ```bash
   docker run -p 8080:8080 go_sample_service:1.0.0
   ```

1. Curl the service (or open in the browser):

   ```bash
   curl http://localhost:8080/httpbin/json
   ```

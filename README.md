# Mock API Server

## Introduction
This app allows you to add urls and serve dummy json responses. It contains two handlers, the `DummyHandler` allows you to map a route to a json schema which will be served as the response and `DummyErrorHandler` allows you to add an endpoint that will return a specific error which is useful for mocking error states.

## Running The Server Locally
Complete the following steps to run this app:

* Ensure golang is installed on your machine
* Build the app using `go build -o bin/hiscox-mock-api`
* Run the app using `./bin/hiscox-mock-api`

The dummy api will now be available on `localhost:8080`

## Adding Routes
Routes can be added to the `router.go` file. Example routes have been added which you can copy. You need only change the url and path to the schema:

```golang
	mux.Handle(
		"/api/test-url",
		&h.DummyHandler{FilePath: "./schemas/you_json_response.json"},
	)
```
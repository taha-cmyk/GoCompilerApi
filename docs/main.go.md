# Go Compiler API
=====================

## Overview
-----------

This API provides a simple interface for compiling Go code. It accepts a POST request with a JSON body containing the code to be compiled and returns the compilation output and any errors that occurred.

## Functions
-------------

### extractOutputAndErrors

#### Description

Extracts the output and errors from a compilation result.

#### Parameters

* `data`: A map of interfaces representing the compilation result.

#### Return Values

* `output`: A string containing the compilation output.
* `errors`: A string containing any compilation errors.

### sendPostRequest

#### Description

Sends a POST request to the Go compiler API with the provided code.

#### Parameters

* `content`: A string containing the code to be compiled.

#### Return Values

* `map[string]interface{}`: A map of interfaces representing the compilation result.
* `error`: An error object if the request failed.

### main

#### Description

Sets up a Gin server to handle compilation requests.

## Usage Examples
-----------------

### Compile Go Code

To compile Go code, send a POST request to the `/compile` endpoint with a JSON body containing the code:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"code": "package main\n\nfunc main() {\n    println(\"Hello, World!\")\n}"}' http://localhost:8080/compile
```

This should return a JSON response with the compilation output and any errors:

```json
{
    "output": "Hello, World!\n",
    "errors": ""
}
```

## API Endpoints
----------------

### /compile

* Method: POST
* Request Body: JSON object with a single field `code` containing the Go code to be compiled.
* Response Body: JSON object with two fields: `output` containing the compilation output and `errors` containing any compilation errors.

## Error Handling
-----------------

The API returns error responses in the following formats:

* `400 Bad Request`: If the request body is invalid or missing required fields.
* `500 Internal Server Error`: If the compilation request fails or the Go compiler API returns an error.

## Dependencies
---------------

* Gin: A high-performance framework for building web applications in Go.
* net/http: The Go standard library for making HTTP requests.
* net/url: The Go standard library for working with URLs.
* encoding/json: The Go standard library for working with JSON data.
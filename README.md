# Go Code Compiler API

This API allows you to compile and run Go code directly from your application. It's designed to be simple and easy to use, making it perfect for educational purposes, code testing, or any other scenario where you need to execute Go code dynamically.

## Endpoints

### POST /compile

This endpoint compiles and runs the provided Go code. It expects a JSON payload with a single key-value pair, where the key is "code" and the value is the Go code to be compiled and executed.

#### Request Body

```json
{
  "code": "fmt.Println(\"Hello, World!\")"
}
```

#### Response

The response will contain two fields: "output" and "errors". "output" will contain the standard output of the executed code, and "errors" will contain any error messages if the compilation or execution fails.

#### Example Response

```json
{
  "output": "Hello, World!\n",
  "errors": ""
}
```

## How to Use

To use this API, you can send a POST request to the `/compile` endpoint with the Go code you want to execute. You can use any HTTP client library or tool to make the request.

Here's an example using `curl` from the command line:

```bash
curl -X POST \
  http://localhost:8080/compile \
  -H 'Content-Type: application/json' \
  -d '{"code": "fmt.Println(\"Hello, World!\")"}'
```

Replace `http://localhost:8080/compile` with the actual URL of the API.

## Error Handling

If there's an error during compilation or execution, the API will return a JSON response with the error message in the "errors" field. The "output" field will be empty in case of an error.

## Note

This API is a wrapper of the official Go online playground,so you can not use some packages like `os/exec` and some other packages that are not allowed in the playground.

## Self-hosting

You can self-host this API by cloning the repository and running `go run main.go` from your terminal , iam not aiming to host this api 

## Contributing 

This API is open-source, and contributions are welcome. If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the GitHub repository.

## License

This API is licensed under the MIT License.







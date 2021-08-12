# HTTP API
This app makes http requests and prints the address of the request along with the
MD5 hash of the response.

It uses GoRoutines to perform requests in parallel, hence, it is possible to use the `-parallel` flag when running it, to limit the concurrent requests preventing local resources exhaustion.

Different URLs can be passed to the app, which will use the native url library to validate if the input is in a valid URL format. 
Even after validation is possible that the given URL is unreachable. This will trigger a log error, printed instead of the expected MD5 response hash. 

## Installation

1. Install Go (see [Download and install](https://golang.org/doc/install))
2. Clone this repository `git clone https://github.com/RaizaClaudino/http-api`
4. Go to the project's folder and run the command below:
    `go build -o myhttp`
5. Use the generated build to run the app

### Sample execution 
```bash
./myhttp -parallel 2 http://www.adjust.com http://google.com facebook.com yahoo.com
```

### Run the tests
```bash
go test ./...
```





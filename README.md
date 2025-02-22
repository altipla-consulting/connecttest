
# connecttest

[![Go Reference](https://pkg.go.dev/badge/github.com/altipla-consulting/connecttest.svg)](https://pkg.go.dev/github.com/altipla-consulting/connecttest)

Helpers to test Connect services.


## Install

```shell
go get github.com/altipla-consulting/connecttest
```


## Usage

```go
reply, err := client.Call(ctx, connect.NewRequest(request))
connecttest.RequireError(t, err, connect.CodeFailedPrecondition, "foo bar baz error message")
```


## Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `make gofmt`.


## License

[MIT License](LICENSE)

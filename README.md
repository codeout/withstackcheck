# withstackcheck

withstackcheck checks error-returning statements are wrapped with `errors.WithStack()` only when the errors come from external packages.

### :o: Valid

```go
func valid() error {
	var err error

	err = throw() // internal
	if err != nil {
		return err
	}

	_, err = json.Marshal(nil) // external
	if err != nil {
		return errors.WithStack(err)
	}
```

### :x: Invalid

```go
func invalid() error {
	var err error

	err = throw() // internal
	if err != nil {
		return errors.WithStack(err) // need to remove errors.WithStack()
	}

	_, err = json.Marshal(nil) // external
	if err != nil {
		return err // need to wrap with errors.WithStack()
	}
```


## Why withstackcheck?

If we want to log an error and stack trace at the outer of the call stack, we need to memorize it at the package boundary.


## Installation

```shell
go install github.com/codeout/withstackcheck/cmd/withstackcheck@latest
```


## How to Use

### go vet

```shell
go vet -vettool=$(go env GOPATH)/bin/withstackcheck ./...
```

### golangci-lint

Not yet supported.


## TODO

### *ast.IndexExpr support

Cannot check indexed return statements:

```go
func unknown() error {
    var errs []error

    if _, err := json.Marshal(nil); err != nil {
    	errs = append(errs, err)
    }

    if len(errs) > 0 {
    	return errs[0] // cannot check this
    }
```


## Copyright and License

Copyright (c) 2023 Shintaro Kojima. Code released under the [MIT license](LICENSE).

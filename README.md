# isbn
ISBN verifier

## Usage

Use `isbnutil <isbn>`. It reports an error and exits with 1, if the ISBN is invalid. Nothing is reported if it's valid.
## Build
```shell
CGOENABLED=0 go build  -o isbnutil -ldflags="-s -w" .
```

Check with `go version -m isbnutil`

# isbn
ISBN verifier

## Usage

Use `isbnutil <isbn>`. It reports an error and exits with 1, if the ISBN is invalid. Nothing is reported if it's valid.
## Build
```shell
go build  -o isbnutil -ldflags="-s -w" .
```

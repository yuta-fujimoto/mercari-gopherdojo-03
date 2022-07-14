# omikuji API

* data resource: https://omikuji-guide.com/number/

## Build

``` 
go build
```

## Test

```
go test ./...
```

## Coverage

```
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```

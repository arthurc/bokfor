# bokfor

## Run tests

```bash
go test -v ./...
```

## Build

```bash
docker build -t bokfor .
```

## Development

```bash
go install github.com/cosmtrek/air@latest
air -c .air.test.toml # To continuously run tests
air # To continuously build the application
```

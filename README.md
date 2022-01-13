# Rescounts API

## Requirements

- Go v1.17
- Optional: Docker

## Development

```bash
$ go mod download
# install required modules

$ go build main.go
# build binary

$ go run main.go
# run main file

$ go test ./... -v
# run all file tests

$ go test -bench=. -benchmem -run=none ./... -v
# run all file benchmarks
```

## Environment Variables

See [config](./internals/config/env.go).

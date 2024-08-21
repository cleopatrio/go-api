OS=linux
ARCH=amd64
GO111MODULE=on
GOENV=GOOS=${OS} GOARCH=${ARCH}

install:
	go mod download -x
	go mod tidy

build-pipe:
	go mod tidy
	env $(GOENV) go build -o $(APP_NAME) -ldflags="-s -w" ./cmd/app/main.go

build-local:
	go mod tidy
	go build -o $(APP_NAME) -ldflags="-s -w" ./cmd/local/main.go

run:
	go run ./cmd/local/main.go

reload_docker:
	./scripts/reload_docker.sh

pprof-local:
	go tool pprof http://localhost:8081/debug/pprof/heap

pprof-rancher:
	go tool pprof http://localhost:8585/debug/pprof/heap

swagger:
	swag fmt
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/app/main.go

wire_server:
	go install github.com/google/wire/cmd/wire
	wire  ./internal/config/injections/server

wire_cli:
	go install github.com/google/wire/cmd/wire
	wire  ./internal/config/injections/cli


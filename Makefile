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
	go run ./cmd/server/main.go

tests:
	go test -v ./test/integration/...
	go test -v ./test/unit/...

docker-up:
	./scripts/docker-up.sh

docker-down:
	./scripts/docker-down.sh


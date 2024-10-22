OS=linux
ARCH=amd64
GO111MODULE=on
GOENV=GOOS=${OS} GOARCH=${ARCH}

scenarios=all

install:
	go mod download -x
	go mod tidy

build-local:
	go mod tidy
	go build -o go-api -ldflags="-s -w" ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

mocks:
	go install go.uber.org/mock/mockgen@latest
	mockgen -source=internal/domain/adapters/create_note_repository_adapter.go -destination=test/mocks/create_note_repository_adapter_mock.go -package=mocks
	mockgen -source=internal/domain/adapters/note_queue_adapter.go -destination=test/mocks/note_queue_adapter_mock.go -package=mocks

tests:
	go test -v ./test/...

bdd:
	go test -v ./test/integration/... --scenarios=$(scenarios)

docker-up:
	./scripts/docker-up.sh

docker-down:
	./scripts/docker-down.sh

install-gremlins:
	go install github.com/go-gremlins/gremlins/cmd/gremlins@main

mutant-test:
	gremlins unleash --config=gremlins.yaml --exclude-files "test/mock/..." --exclude-files "test/mocks/..."
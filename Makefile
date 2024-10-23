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
	go install go.uber.org/mock/mockgen
	mockgen -source=internal/domain/adapters/create_note_repository_adapter.go -destination=test/mocks/create_note_repository_adapter_mock.go -package=mocks
	mockgen -source=internal/domain/adapters/create_note_repository_adapter.go -destination=test/mocks/create_note_repository_adapter_mock.go -package=mocks
	mockgen -source=internal/integration/queues/notes_queue.go -destination=test/mocks/notes_queue_mock.go -package=mocks


tests:
	go test -v ./test/...

bdd:
	go test -v ./test/integration/... --scenarios=$(scenarios)

docker-up:
	./scripts/docker-up.sh

docker-down:
	./scripts/docker-down.sh

mutant-test:
	go get github.com/go-gremlins/gremlins/cmd/gremlins
	go install github.com/go-gremlins/gremlins/cmd/gremlins
	gremlins unleash --config=gremlins.yaml --exclude-files "test/mock/..." --exclude-files "test/mocks/..."
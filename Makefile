EXEC_PATH = ./build/dist
BUILD_CMD = go build -o $(EXEC_PATH) -ldflags="-s -w"

lint:
	golangci-lint run

dev:
	air

build:
	$(BUILD_CMD) ./cmd/cli

run: build
	$(EXEC_PATH) new

build-web:
	$(BUILD_CMD) ./cmd/server

run-web: build-web
	$(EXEC_PATH)

web:
	cd client && npm run start

test:
	go test -json  -count=1 ./... -coverpkg=./... -coverprofile coverage.out -covermode=atomic | gotestfmt && go tool cover -html coverage.out && rm coverage.out

clean:
	rm -rf ./build/

.PHONY: lint run dev run-web test build build-web clean web

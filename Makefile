EXECPATH = ./build/dist

lint:
	golangci-lint run

dev:
	air

build:
	go build -o $(EXECPATH) .

run: build
	$(EXECPATH) new

run-web:
	go run main.go --web=t

web:
	cd client && npm run start

test:
	go test -json  -count=1 ./... -coverpkg=./... -coverprofile coverage.out -covermode=atomic | gotestfmt # && go tool cover -html coverage.out && rm coverage.out

.PHONY: lint run dev run-web test build

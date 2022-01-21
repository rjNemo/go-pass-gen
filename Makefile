lint:
	golangci-lint run

dev:
	air

run:
	go run main.go

run-web:
	go run main.go --web=t

web:
	cd client && npm run start

test:
	go test -json  -count=1 ./... -coverpkg=./... -coverprofile coverage.txt -covermode=atomic | gotestfmt && go tool cover -html coverage.txt && rm coverage.txt

.PHONY: lint run dev run-web test
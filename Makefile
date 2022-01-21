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

.PHONY: lint run dev run-web
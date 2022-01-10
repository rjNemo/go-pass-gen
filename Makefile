lint:
	golint cmd/...
	golint passgen/...

dev:
	air

run:
	go main.go

run-web:
	go main.go --web=t

web:
	cd client && npm run start

.PHONY: lint run dev run-web
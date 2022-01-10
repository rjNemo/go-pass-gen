lint:
	golint cmd/...
	golint passgen/...

run:
	air

web:
	cd client && npm run start

.PHONY: lint run
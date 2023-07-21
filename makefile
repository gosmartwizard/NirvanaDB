tidy:
	go mod tidy
	go mod vendor

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

run:
	go run main.go

build:
	go build -o /home/nk0379/go/bin/ndb main.go
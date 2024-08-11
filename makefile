run:
	go run cmd/main.go

fmt:
	go fmt ./...

lint:
	go vet ./...

ts: # test
	go test -v ./...

build:
	go build cmd/main.go

release:
	go build -ldflags "-s -w" cmd/main.go
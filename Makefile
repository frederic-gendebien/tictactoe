PROJECT=tictactoe
VERSION=$$(git rev-parse --short=10 HEAD)

run:
	go run cmd/main.go

bin:
	mkdir bin

bin/app: bin
	go build -o bin/app cmd/main.go

container: bin/app
	docker build -f build/Dockerfile . -t $(PROJECT):$(VERSION)
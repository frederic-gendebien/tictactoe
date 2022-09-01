PROJECT=tictactoe
VERSION=$$(git rev-parse --short=10 HEAD)

run:
	go run cmd/main.go

container:
	docker build -f build/Dockerfile . -t $(PROJECT):$(VERSION)

container-run: container
	docker run -it --rm $(PROJECT):$(VERSION)
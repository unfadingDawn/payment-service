lint:
	golangci-lint run --fix
up:
	sudo docker compose up -d
	go build
test:
	go test -v -race -count=1 ./...

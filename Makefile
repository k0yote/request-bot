build:
	go build -o ./bin/simplebot

run: build
	./bin/simplebot

test:
	go test -v ./...

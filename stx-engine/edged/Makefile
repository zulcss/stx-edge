clean:
	rm -rf bin

client:
	go build -o bin/stx cmd/client/main.go

server:
	go build -o bin/edged cmd/server/main.go

all: clean client server

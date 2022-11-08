build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o directexchange ./...

run: directexchange
	./directexchange

clean: directexchange
	@rm ./directexchange

all: build run clean
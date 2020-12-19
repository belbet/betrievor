binary = betrievor

all: linux
	
linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)
	GOOS=linux GOARCH=amd64 go build -o ./$(binary)
macos:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(binary)
	GOOS=darwin GOARCH=amd64 go build -o ./$(binary)

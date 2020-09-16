build:
	go build -o bin server.go
run:
	go run server.go
compile:
	GOOS=darwin GOARCH=amd64 go build -o bin/server-darwin-amd64 server.go
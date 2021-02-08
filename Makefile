build: build-macos build-linux build-windows

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o azure-sql-test-darwin-amd64

build-linux:
	GOOS=linux GOARCH=amd64 go build -o azure-sql-test-linux-amd64

build-windows:
	GOOS=windows GOARCH=amd64 go build
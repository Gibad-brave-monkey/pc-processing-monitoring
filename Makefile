include .env

build_for_windows:
	GOOS=windows GOARCH=amd64 go build -o ${FILE_NAME_WINDOWS}.exe cmd/process/main.go

build_for_macOS:
	GOOS=darwin GOARCH=amd64 go build -o ${FILE_NAME_MACOS} cmd/process/main.go

build_for_linux:
	GOOS=linux GOARCH=amd64 go build -o ${FILE_NAME_LINUX} cmd/process/main.go

run:
	go run cmd/process/main.go

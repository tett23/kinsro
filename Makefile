
test:
	@GO_ENV=test go test ./... -timeout 1s

test-v:
	@GO_ENV=test go test ./... -timeout 1s -v

build:
	@GOOS=linux GOARCH=amd64 go build -o build/kinsro_linux_amd64 src/main.go
	@GOOS=linux GOARCH=arm go build -o build/kinsro_linux_arm32	src/main.go
	@GOOS=linux GOARCH=arm64 go build -o build/kinsro_linux_arm64 src/main.go
	@GOOS=darwin GOARCH=amd64 go build -o build/kinsro_darwin_amd64 src/main.go
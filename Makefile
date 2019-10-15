
test:
	@GO_ENV=test go test ./... -timeout 1s

test-v:
	@GO_ENV=test go test ./... -timeout 1s -v

build: src/main.go
	@GOOS=linux GOARCH=amd64 go build -o build/kinsro_linux_amd64 src/main.go
	@GOOS=linux GOARCH=arm go build -o build/kinsro_linux_arm32	src/main.go
	@GOOS=linux GOARCH=arm64 go build -o build/kinsro_linux_arm64 src/main.go
	@GOOS=darwin GOARCH=amd64 go build -o build/kinsro_darwin_amd64 src/main.go
	@GOOS=js GOARCH=wasm go build -o front/src/kinsro.wasm src/vindex/vindexdata/*.go
SOURCES := $(shell find . -type f -name '*.go')
FRONT_SOURCES := $(shell find ./front/src -type f -name '*.ts' -or -name '*.tsx' )

test:
	@GO_ENV=test go test ./... -timeout 1s

test-v:
	@GO_ENV=test go test ./... -timeout 1s -v

build: $(SOURCES) 
	@GOOS=linux GOARCH=amd64 go build -o build/kinsro_linux_amd64 src/main.go
	@GOOS=linux GOARCH=arm go build -o build/kinsro_linux_arm32	src/main.go
	@GOOS=linux GOARCH=arm64 go build -o build/kinsro_linux_arm64 src/main.go
	@GOOS=darwin GOARCH=amd64 go build -o build/kinsro_darwin_amd64 src/main.go
	@GOOS=js GOARCH=wasm go build -o front/public/kinsro.wasm src/js_main/js_main.go
	@GOOS=js GOARCH=wasm go build -o build/front/kinsro.wasm src/js_main/js_main.go

build-front: $(FRONT_SOURCES)
	@yarn run build

deploy: build build-front
	scripts/deploy
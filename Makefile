.PHONY: all clean build serve test

BINARY=main.wasm

all: build serve

clean:
	rm -Rf "$(BINARY)"

build:
	GOOS=js GOARCH=wasm go build -o $(BINARY) 

serve:
	go run dev.go

test:
	GOOS=js GOARCH=wasm go test -v -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec"
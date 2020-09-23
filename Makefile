GC=go

PROJECT=kokos_quest

MAIN_PKG=kokos_quest/cmd/kokos_quest
WASM_SERVER_PKG=kokos_quest/cmd/wasm_server
EDITOR_PKG=kokos_quest/cmd/editor

OUTPUT_DIR=./build
WASM_DIR=$(OUTPUT_DIR)/html

ifeq ($(GOOS), windows)
	EXE := .exe
endif

ifdef COMSPEC
	EXE := .exe
endif

GLOBAL_ARGS=
OUTPUT=$(OUTPUT_DIR)/$(PROJECT)$(EXE)
WASM_OUTPUT=$(WASM_DIR)/assets/wasm/$(PROJECT).wasm

.PHONY: % run editor build fmt clean generate serve wasm

%:
	@:

# args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

all: build

editor:
	@$(GC) run $(EDITOR_PKG) $(args)

run:
	@$(GC) run $(MAIN_PKG)

build:
	@echo "Compiling to $(OUTPUT)"
	@$(GC) build $(GLOBAL_ARGS) -o $(OUTPUT) $(MAIN_PKG)
	@echo "Done!"

fmt:
	@$(GC) fmt ./...

clean:
	@rm $(OUTPUT_DIR) -r

gen:
	@echo "Generating all..."
	@$(GC) generate -v ./...
	@echo "Done!"

wasm:
	@echo "Copying markup to $(WASM_DIR)"
	@mkdir $(WASM_DIR) -p
	@cp -r html/* $(OUTPUT_DIR)/html
	@echo "Compiling to $(WASM_OUTPUT)"
	@GOOS=js GOARCH=wasm $(GC) build $(GLOBAL_ARGS) -ldflags "-s -w" -o $(WASM_OUTPUT) $(MAIN_PKG)
	@echo "Gziping..."
	@gzip -9 -v -c $(WASM_OUTPUT) > $(WASM_OUTPUT).gz
#	@echo "Optimizing..."
#	@wasm-opt $(WASM_OUTPUT) -O3 -o $(WASM_OUTPUT)
#	@echo "Striping..."
#	@wasm-strip $(WASM_OUTPUT)
	@echo "Copying glue JS code to $(WASM_DIR)/assets/js/wasm_exec.js"
	@cp `$(GC) env GOROOT`/misc/wasm/wasm_exec.js $(WASM_DIR)/assets/js
	@echo "Done!"

wasm-server:
	@$(GC) run $(WASM_SERVER_PKG) -gzip -path $(WASM_DIR)

serve: wasm wasm-server

serve-live:
	@echo "Serving to :8080 port"
	@wasmserve $(MAIN_PKG)

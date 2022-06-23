.DEFAULT_GOAL: build

compile_qtpl:
	go run view/qtpl_compiler/main.go || true
.PHONY: complie_qtpl

run: compile_qtpl
	go run main.go || true
.PHONY: run

build: compile_qtpl
	go build main.go || true
.PHONY: build
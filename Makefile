.DEFAULT_GOAL: build

complie_qtpl:
	go run view/qtc/main.go || true
.PHONY: complie_qtpl

run: complie_qtpl
	go run main.go || true
.PHONY: run

build: complie_qtpl
	go build main.go || true
.PHONY: build
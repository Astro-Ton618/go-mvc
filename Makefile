.DEFAULT_GOAL: build

complie_qtpl:
	go run view/qtc/main.go
.PHONY: complie_qtpl

run: complie_qtpl
	go run main.go
.PHONY: run

build: complie_qtpl
	go build main.go
.PHONY: build
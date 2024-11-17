.DEFAULT_GOAL=run


.PHONY: run

run-engine:
	go run ./cmd/engine/engine.go

run:
	go run ./cmd/desktop/Desktop.go

build:
	go build -o app ./cmd/desktop/Desktop.go

test:
	go test ./... -v

fmt:
	@echo "TODO"

debug:
	go run -tags debug ./cmd/desktop/Desktop.go

clear:
	@rm app
.DEFAULT_GOAL=run


.PHONY: run

run:
	go run ./cmd/desktop/Desktop.go

build:
	go build -o app ./cmd/desktop/Desktop.go

fmt:
	@echo "TODO"

debug:
	go run -tags debug ./cmd/desktop/Desktop.go

clear:
	@rm app
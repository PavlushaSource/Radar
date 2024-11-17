.DEFAULT_GOAL=run


.PHONY: run

run:
	go run ./cmd/Desktop.go

build:
	go build -o app ./cmd/Desktop.go

fmt:
	@echo "TODO"

debug:
	go run -tags debug ./cmd/Desktop.go

clear:
	@rm task
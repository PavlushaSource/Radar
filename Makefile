.DEFAULT_GOAL=run


.PHONY: run

run-engine:
	go run ./cmd/engine/engine.go

run:
	go run ./cmd/desktop/desktop.go

build:
	go build -o app ./cmd/desktop/Desktop.go

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

clear:
	@rm app
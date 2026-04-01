GOCMD=go
APP=logreport

.PHONY: fmt test build run

fmt:
	$(GOCMD) fmt ./...

test:
	$(GOCMD) test ./...

build:
	$(GOCMD) build -o bin/$(APP) ./cmd/logreport

run:
	$(GOCMD) run ./cmd/logreport -in ./examples/sample.log

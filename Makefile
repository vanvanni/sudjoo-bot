.PHONY: help
help:
	@echo [
	@echo [	Use one of the options: tidy / run / dev
	@echo [

# App
MAIN=cmd/imgdn/main.go
AIR=air
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOMODTIDY=$(GOCMD) mod tidy

.PHONY: tidy run dev
tidy:
	$(GOMODTIDY)

run:
	$(GOCMD) run $(MAIN)

dev:
	$(AIR)
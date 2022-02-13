NAME=jasoncoding-backendv2

all: help

## build: Compiles the app.
.PHONY: build
build:
	@go build -o $(NAME)

build:
.PHONY: run
run:
	@./$(NAME)

.PHONY: run-prod
run-prod:
	@ENVIRONMENT=production ./$(NAME)

.PHONY: deps
deps:
	@go mod download

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
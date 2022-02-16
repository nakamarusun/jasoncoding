NAME=jcbackendv2

all: help

.PHONY: build
## build: Compiles the app.
build:
	@go build -o $(NAME)

build:
.PHONY: run
## run: Runs the app in dev mode.
run:
	@./$(NAME)

.PHONY: run-prod
## run-prod: Runs the app in production mode.
run-prod:
	@ENVIRONMENT=production ./$(NAME)

.PHONY: deps
## deps: Install dependencies.
deps:
	@go mod download

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
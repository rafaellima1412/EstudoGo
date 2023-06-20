.PHONY: default run build test docs clean
APP_NAME = estudoGo

default: run-with-docs

run:
	@go run  main.go
run-with-docs:
	@swag init
	@go run  main.go
build:
	@go build -o $(APP_NAME) main.go
teste:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
cache:
	@go clean -cache
include .env.example
export $(shell sed 's/=.*//' .env.example)

dev:
	reflex -r '(\.go$$|go\.mod)' -s go run .

build:
	docker build -t

test:
	@echo "== 🦸‍️ ci.tester =="
	ginkgo -r -v

ci.lint:
	@echo "== 🙆 ci.linter =="
	golangci-lint run -v ./... --fix

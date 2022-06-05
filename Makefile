init:
	@echo "== 👩‍🌾 init =="

	brew install go
	brew install node
	brew install pre-commit
	brew install golangci-lint
	brew upgrade golangci-lint

c.test:
	@echo "== 🦸‍♀️ ci.tester =="

ci.lint:
	@echo "== 🦸‍♀️ ci.linter =="
	golangci-lint run -v ./... --fix

init:
	@echo "== 👩‍🌾 init =="
	brew install go
	brew install node
	brew install pre-commit
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install

	@echo "== install ginkgo =="
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

	@echo "== install gomock =="
	go install github.com/golang/mock/mockgen@v1.6.0

precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

test:
	@echo "== 🦸‍️ ci.tester =="

ci.lint:
	@echo "== 🙆 ci.linter =="
	golangci-lint run -v ./... --fix

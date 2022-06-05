init:
    @echo "== init =="
    brew install go
    brew install node
    brew install pre-commit

run:dev:
	docker compose up --build.

greeting:
	@echo "hello".

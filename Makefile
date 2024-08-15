.PHONY:
setup:
	docker compose build --no-cache
	cd web && [ -f .env ] || cp .env.example .env && npm ci

.PHONY:
up:
	docker compose up -d
	cd web && npm run dev

.PHONY:
down:
	docker compose down

.PHONY:
oapi:
	docker compose run --rm api bash -c "cd internal/pkg && oapi-codegen -package oapi ../../api/api.yaml > ./oapi/server.go"
	go mod tidy

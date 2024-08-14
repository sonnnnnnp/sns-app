GO_TO_BACKEND = cd packages/backend
GO_TO_FRONTEND = cd packages/frontend

.PHONY:
setup:
	docker compose build --no-cache
	$(GO_TO_FRONTEND) && [ -f .env ] || cp .env.example .env && npm ci

.PHONY:
up:
	docker compose up -d
	$(GO_TO_FRONTEND) && npm run dev

.PHONY:
down:
	docker compose down

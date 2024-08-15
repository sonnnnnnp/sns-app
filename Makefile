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

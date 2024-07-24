GO_TO_BACKEND = cd packages/backend
GO_TO_FRONTEND = cd packages/frontend

.PHONY:
setup:
	docker compose build --no-cache
	$(GO_TO_BACKEND) && [ -f .env ] || cp .env.example .env && npm ci
	$(GO_TO_FRONTEND) && [ -f .env ] || cp .env.example .env && npm ci

.PHONY:
up:
	docker compose up -d
	@$(GO_TO_BACKEND) && npm run start:dev & \
	$(GO_TO_FRONTEND) && npm run dev

.PHONY:
down:
	docker compose down

.PHONY:
seed:
	$(GO_TO_BACKEND) && npx prisma db push
	$(GO_TO_BACKEND) && npx prisma db seed

.PHONY:
migrate:
	$(GO_TO_BACKEND) && npx prisma migrate dev --name ${name}

.PHONY:
view-db:
	$(GO_TO_BACKEND) && npx prisma studio

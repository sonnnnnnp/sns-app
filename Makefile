.PHONY: help setup up down openapi

#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"

#? setup: アプリケーションのセットアップ
setup:
	docker compose build --no-cache
	cd web && [ -f .env ] || cp .env.example .env && npm ci

#? up: アプリケーションの起動
up:
	docker compose up -d
	cd web && npm run dev

#? down: アプリケーションの停止
down:
	docker compose down

#? openapi: OpenAPI からコードを生成
openapi:
	docker compose run --rm api bash -c "cd internal/pkg && oapi-codegen -package oapi ../../api/openapi.yaml > ./oapi/server.go"

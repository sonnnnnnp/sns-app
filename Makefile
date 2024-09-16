.PHONY: help setup up down tidy

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

#? tidy: Go モジュールの依存関係を整理
tidy:
	docker compose run --rm api bash -c "go mod tidy"

.PHONY: openapi ent-new ent-gen

#? oapi: OpenAPI からコードを生成
oapi:
	docker compose run --rm api bash -c "cd pkg && oapi-codegen -package oapi ../api/openapi.yaml > ./oapi/server.go"
	npx openapi-typescript ./api/openapi.yaml -o ./web/lib/api/client.ts

#? ent-new: ent エンティティを生成
ent-new:
	docker compose run --rm api bash -c "cd ./internal/domain && go run -mod=mod entgo.io/ent/cmd/ent new $(name)"

#? ent-gen: ent エンティティからコードを生成
ent-gen:
	docker compose run --rm api bash -c "go generate ./internal/domain/ent"

#? wire: 依存関係の自動生成
wire:
	docker compose run --rm api bash -c "cd ./internal && wire gen"

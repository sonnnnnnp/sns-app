.PHONY: help build up down shell tidy

#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"

#? build: アプリケーションのセットアップ
build:
	docker compose build --no-cache
	[ -f .env ] || cp .env.local .env
	npm ci && cd web && npm ci

#? up: アプリケーションの起動
up:
	docker compose up -d
	cd web && npm run dev

#? down: アプリケーションの停止
down:
	docker compose down

#? shell: API コンテナのシェルを起動
shell:
	docker compose exec -it httpserver bash

#? tidy: Go モジュールの依存関係を整理
tidy:
	docker compose run --rm httpserver bash -c "go mod tidy"

.PHONY: api wire

#? api: OpenAPI からコードを生成
api:
	docker compose run --rm httpserver bash -c "cd pkg && oapi-codegen -package api /httpserver/api/openapi.json > /httpserver/internal/server/http/adapter/api/interface.go"
	npx openapi-typescript ./api/openapi.json -o ./web/app/api/client.ts

#? wire: サーバーの依存関係を自動生成
wire:
	docker compose run --rm httpserver bash -c "cd /httpserver/pkg/httpserver/http && wire gen && mv ./wire_gen.go /httpserver/internal/server/http/wire.go"

.PHONY: migrate migrate-down seed sql-gen

#? migrate: データベースの構造をマイグレート
migrate:
	docker compose run --rm httpserver bash -c "migrate -source file://pkg/db/migrations -database postgres://user:password@db:5432/db?sslmode=disable up"

#? migrate-down: データベースの構造を初期化
migrate-down:
	docker compose run --rm httpserver bash -c "migrate -source file://pkg/db/migrations -database postgres://user:password@db:5432/db?sslmode=disable down"

#? seed: データベースへ初期データを投入
seed:
	docker compose run --rm httpserver bash -c "go run ./cmd/seeder"

#? sql-gen: SQL クエリから Go コードを生成
sql-gen:
	docker compose run --rm httpserver bash -c "sqlc generate"

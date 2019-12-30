.DEFAULT_GOAL := help

.PHONY: serve
serve:  ## 開発サーバを起動する
	docker-compose up -d app

.PHONY: down
down:  ## 開発サーバを停止する
	docker-compose down

.PHONY: logs
logs:  ## アプリケーションログの表示
	docker-compose logs -f app

.PHONY: cli
cli:  ## gRPCクライアントの起動
	docker-compose run --rm evans

.PHONY: protoc
protoc:  ## codegenの実行
	docker-compose run --rm protoc

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

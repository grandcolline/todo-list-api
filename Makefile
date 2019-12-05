.PHONY: serve
serve:
	docker-compose up -d app

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f app

.PHONY: cli
cli:
	docker-compose run --rm evans

.PHONY: protoc
protoc:
	docker-compose run --rm protoc

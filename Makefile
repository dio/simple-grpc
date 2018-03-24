rpc=hello

up: down
	docker-compose up --build --force-recreate  --always-recreate-deps

down:
	docker-compose down

build: api
	docker-compose build --no-cache

call:
	@docker-compose exec client $(rpc)

.PHONY: build up down api call
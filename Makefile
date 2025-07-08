.PHONY: start down restart migrate

start:
	docker compose up --build -d

down:
	docker compose down --volumes

restart:
	docker compose restart

migrate:
	docker compose run --rm migrate

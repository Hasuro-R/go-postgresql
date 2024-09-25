.PHONY: compose/build compose/run compose/down

compose/build:
	docker-compose build

compose/up:
	docker-compose up

compose/down:
	docker-compose down

.PHONY: db/migrate db/migrate/down

db/migrate:
	docker compose run api sql-migrate up

db/migrate/down:
	docker compose run api sql-migrate down

g/migrate:
	docker compose run api sql-migrate new ${name}

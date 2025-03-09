.PHONY: run
run:
	go run ./cmd/server/main.go

.PHONY: db
db:
	docker compose up -d


.PHONY: migrate
migrate:
	go run ./cmd/migrate/main.go -migrationsFolder migrations


.PHONY: clean
clean:
	docker compose down && docker volume rm vyking_app_mysql_data

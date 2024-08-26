install-local:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	templ generate

tailwind-watch:
	tailwindcss -i ./pkg/templates/tailwind.css -o ./public/css/main.css --watch

tailwind-build:
	tailwindcss -i ./pkg/templates/tailwind.css -o ./public/css/main.css --minify

up-database:
	docker compose up -d --force-recreate database

db-connect:
	docker exec -it homecomp-database-1 mariadb -u $(HCMP_DB_USER) -p$(HCMP_DB_PASS)

migrate-up:
	migrate -database "mysql://$(HCMP_DB_USER):$(HCMP_DB_PASS)@tcp($(HCMP_DB_HOST):$(HCMP_DB_PORT))/$(HCMP_DB_NAME)" -path ./migrations up

migrate-down:
	migrate -database "mysql://$(HCMP_DB_USER):$(HCMP_DB_PASS)@tcp($(HCMP_DB_HOST):$(HCMP_DB_PORT))/$(HCMP_DB_NAME)" -path ./migrations down

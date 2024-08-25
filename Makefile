local-environment:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	templ generate

templ-gen:
	templ generate

app-watch:
	air

tailwind-watch:
	tailwindcss -i ./pkg/templates/tailwind.css -o ./public/css/main.css --watch

tailwind-build:
	tailwindcss -i ./pkg/templates/tailwind.css -o ./public/css/main.css --minify

up-database:
	docker compose up -d --force-recreate database

db-connect:
	docker exec -it homecomp-database-1 mariadb -u $(HCMP_DB_USER) -p$(HCMP_DB_PASS)

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

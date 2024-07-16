local-environment:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest

templ-gen:
	templ generate

tailwind-watch:
	tailwind -i ./pkg/templates/tailwind.css -o ./public/css/main.css --watch

tailwind-build:
	tailwind -i ./pkg/templates/tailwind.css -o ./public/css/main.css --minify

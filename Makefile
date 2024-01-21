install-binaries:
	go install github.com/mitranim/gow@latest

run-watch:
	gow -e=go,mod,html -c run main.go

run-prod:
	ENV=_prod go run main.go
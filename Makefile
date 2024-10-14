go-build:
	go run main.go

run-db-local:
	cd component/database && docker compose up

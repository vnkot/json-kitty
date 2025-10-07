up-dev:
	go run ./cmd/main.go

up-prod:
	docker compose up -d

down-prod:
	docker compose down
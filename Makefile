up-dev:
	docker compose -f ./docker-compose.dev.yml up -d

down-dev:
	docker compose -f ./docker-compose.dev.yml down

up-prod:
	docker compose up -d

down-prod:
	docker compose down
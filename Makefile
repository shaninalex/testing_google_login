start:
	docker compose --env-file .env up -d --build

down:
	docker compose --env-file .env down
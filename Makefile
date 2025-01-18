infra-up:
	docker compose -f ./local/docker-compose.yml -p warehouse  up -d

infra-down:
	docker compose -f ./local/docker-compose.yml -p warehouse down

infra-reset:
	docker compose -f ./local/docker-compose.yml -p warehouse down
	rm -rf local/data/*

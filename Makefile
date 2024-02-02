include .env

docker_postgres:
	docker run --name=$(DB_NAME) --rm -d -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -p 5432:5432 postgres
migrate_postgres:
	migrate -path ./schema -database '$(DBPG_PATH)' up
docker_nats_streaming:
	docker run --name wb_orders --rm -d -p 4222:4222 -p 8222:8222 -p 6222:6222 nats-streaming -cid $(CLUSTER_ID) -mc 1 -msu 1
sender:
	go run cmd/order_sender/order_sender.go 
app:
	go run cmd/app/main.go
up:
	docker compose up --build
down:
	docker compose down

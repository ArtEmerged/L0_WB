include wb.env

docker_postgres:
	docker run --name=$(DB_NAME) --rm -d -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -e POSTGRES_DB=$(DB_NAME) -p 5432:5432 postgres
migrate_postgres:
	migrate -path ./schema -database 'postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)' up
docker_nats_streaming:
	docker run --name wb_orders --rm -d -p 4222:4222 -p 8222:8222 -p 6222:6222 nats-streaming -cid wb_orders -mc 1 -msu 1
sender:
	go run cmd/order_sender/order_sender.go 
app:
	go run cmd/app/main.go
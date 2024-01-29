include wb.env

docker_postgres:
	docker run --name=$(DB_NAME) --rm -d -e DB_PASSWORD=$(DB_PASSWORD) -v $(HOST_DATA_DIR):/var/lib/postgresql/data -p 5432:5432 postgres
docker_nats_streaming:
	docker run --name wb_orders --rm -d -p 4222:4222 -p 8222:8222 -p 6222:6222 nats-streaming -cid wb_orders -mc 1 -msu 1
sender:
	go run cmd/order_sender/order_sender.go 
app:
	go run cmd/app/main.go
docker_nats_streaming:
	docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 nats-streaming -cid wb_orders
sender:
	go run cmd/order_sender/order_sender.go 
app:
	go run cmd/app/main.go
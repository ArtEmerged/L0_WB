docker_nats_streaming:
	docker run  -p 4222:4222 -p 8222:8222 -p 6222:6222 --name wb_orders nats-streaming -cid wb_orders -mc 1 -msu 1
sender:
	go run cmd/order_sender/order_sender.go 
app:
	go run cmd/app/main.go
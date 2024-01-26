package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "wb_orders"
	clientID  = "order_recipient"
	channel   = "order"
	durableID = "wb-durable"
)

func main() {
	sc, err := stan.Connect(
		clusterID,
		clientID)
	if err != nil {
		log.Printf("[ERROR]: %s", err.Error())
		return
	}
	defer sc.Close()
	for i := 0; i < 10; i++ {
		sub, err := sc.Subscribe(channel, func(msg *stan.Msg) {
			data := string(msg.Data)
			fmt.Println(data)
			msg.Ack()
		},
			stan.DurableName(durableID))
		if err != nil {
			log.Printf("[ERROR]: %s", err.Error())
			return
		}
		time.Sleep(5 * time.Second)
		sub.Close()
	}
}

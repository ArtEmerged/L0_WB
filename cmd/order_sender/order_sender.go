package main

import (
	"fmt"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
)

const (
	clusterID = "wb_orders"
	clientID  = "order_sender"
	channel   = "order"
)

func main() {
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Printf("[ERROR]: %s", err.Error())
		return
	}
	defer sc.Close()

	for i := 0; ; i++ {
		msg := fmt.Sprintf("Hello %d", i)
		err := sc.Publish(channel, []byte(msg))
		if err != nil {
			log.Printf("[ERROR]: %s", err.Error())
			return
		}
		fmt.Println("I send:" + msg)
		time.Sleep(5 * time.Second)
	}
}

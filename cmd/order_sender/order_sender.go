package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"wblzero/internal/models"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
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
		logrus.Error(err.Error())
		return
	}
	defer sc.Close()

	var order models.Order
	ticker := time.NewTicker(time.Second * 5)
	logrus.Infof("connected to Nats Streaming %s clusterID: [%s] clientID: [%s]", stan.DefaultNatsURL, clientID, clusterID)
	for range ticker.C {
		order = randomOrder()
		orderByte, err := json.Marshal(order)
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		err = sc.Publish(channel, orderByte)
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		logrus.Infof("order:%s successfully shipped to '%s' channel", order.OrderUID, channel)
	}
}

func randomOrder() models.Order {
	randomOrder := models.Order{
		OrderUID:    generateRandomString(20),
		TrackNumber: generateRandomString(12),
		Entry:       generateRandomString(4),
		Delivery: models.Delivery{
			Name:    generateRandomString(10),
			Phone:   generateRandomPhoneNumber(),
			ZIP:     generateRandomString(7),
			City:    generateRandomString(15),
			Address: generateRandomString(20),
			Region:  generateRandomString(10),
			Email:   generateRandomEmail(),
		},
		Payment: models.Payment{
			Transaction:  generateRandomString(20),
			RequestID:    generateRandomString(10),
			Currency:     "USD",
			Provider:     generateRandomString(6),
			Amount:       uint(rand.Intn(2000) + 1000),
			PaymentDt:    uint(time.Now().Unix()),
			Bank:         generateRandomString(5),
			DeliveryCost: uint(rand.Intn(500) + 500),
			GoodsTotal:   uint(rand.Intn(500) + 500),
			CustomFee:    uint(rand.Intn(100)),
		},
		Items: []models.Item{
			{
				ChrtID:      uint(rand.Intn(100000)),
				TrackNumber: generateRandomString(12),
				Price:       uint(rand.Intn(500) + 500),
				Rid:         generateRandomString(20),
				Name:        generateRandomString(8),
				Sale:        uint8(rand.Intn(50)),
				Size:        generateRandomString(2),
				TotalPrice:  uint(rand.Intn(500) + 500),
				NmID:        uint(rand.Intn(100000)),
				Brand:       generateRandomString(12),
				Status:      uint16(rand.Intn(500)),
			},
		},
		Locale:            "en",
		InternalSignature: generateRandomString(15),
		CustomerID:        generateRandomString(5),
		DeliveryService:   generateRandomString(5),
		Shardkey:          generateRandomString(2),
		SmID:              uint(rand.Intn(100)),
		DateCreated:       time.Now().UTC(),
		OofShard:          generateRandomString(1),
	}

	return randomOrder
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func generateRandomPhoneNumber() string {
	return fmt.Sprintf("+7%d%d", 1+rand.Intn(999), rand.Intn(10000000))
}

func generateRandomEmail() string {
	return fmt.Sprintf("user%d@example.com", rand.Intn(1000))
}

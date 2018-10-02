package main

import (
	"github.com/bank-now/bn-common-io/queues/pub"
	"log"
)

func main() {
	c := pub.Config{Address: "192.168.88.24:4150",
		Version: "v1",
		Name:    "manager",
		Topic:   "interest-calculation-v1"}

	producer, err := pub.Setup(c)
	if err != nil {
		log.Fatal(err)
	}
	producer.Stop()

}

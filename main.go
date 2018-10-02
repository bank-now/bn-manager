package main

import (
	"fmt"
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

	for i := 1; i <= 10; i++ {
		producer.Publish(c.Topic, []byte(fmt.Sprint(i)))
	}

	producer.Stop()

}

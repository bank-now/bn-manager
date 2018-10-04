package main

import (
	"fmt"
	"github.com/bank-now/bn-common-io/queues/pub"
	"github.com/bank-now/bn-common-model/common/operation"
	"log"
	"time"
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

		item := operation.InterestOperationV1{
			Account: fmt.Sprint(i),
			DateFor: time.Now()}
		b, _ := item.ToJsonBytes()
		producer.Publish(c.Topic, b)
	}

	producer.Stop()

}

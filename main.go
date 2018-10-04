package main

import (
	"fmt"
	"github.com/bank-now/bn-common-io/queues/pub"
	"github.com/bank-now/bn-common-model/common/operation"
	"log"
)

const (
	version = "v1"
	name    = "manager"
)

func main() {
	c := pub.Config{Address: "192.168.88.24:4150",
		Name:    name,
		Version: version,
		Topic:   operation.InterestOperationV1Topic}

	producer, err := pub.Setup(c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		item := operation.NewInterestOperationV1(fmt.Sprint(i))
		b, _ := item.ToJsonBytes()
		producer.Publish(c.Topic, b)
	}
	producer.Stop()

}

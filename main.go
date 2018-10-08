package main

import (
	"fmt"
	"github.com/bank-now/bn-common-io/queues/pub"
	"github.com/bank-now/bn-common-io/zipkin"
	"github.com/bank-now/bn-common-model/common/operation"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

const (
	version   = "v1"
	name      = "manager"
	Address   = "192.168.88.24:4150"
	ZipKinUrl = "http://192.168.88.24:9411/api/v2/spans"
)

var (
	FullName = fmt.Sprint(name, "-", version)
	action   = fmt.Sprint(operation.InterestOperationV2Topic, ".", "createItem")
)

func main() {
	c := pub.Config{Address: Address,
		Name:    name,
		Version: version,
		Topic:   operation.InterestOperationV2Topic}

	producer, err := pub.Setup(c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		ghost := addOneWorkItem(producer, c, fmt.Sprint(i))
		item := operation.NewInterestOperationV2(fmt.Sprint(i))
		item.Ghost = ghost
		b, _ := item.ToJsonBytes()
		producer.Publish(c.Topic, b)
	}
	producer.Stop()

}

func addOneWorkItem(producer *nsq.Producer, c pub.Config, acc string) zipkin.Ghost {
	start := time.Now()
	item := operation.NewInterestOperationV2(fmt.Sprint(acc))
	b, _ := item.ToJsonBytes()
	producer.Publish(c.Topic, b)

	ns := time.Since(start).Nanoseconds()
	ghost := zipkin.LogParent(ZipKinUrl, FullName, action, ns)
	return ghost

}

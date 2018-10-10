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
	serviceName = "InterestCalculation"
	Version     = "v1"
	Name        = "manager"
	Address     = "192.168.88.24:4150"
	ZipKinUrl   = "http://192.168.88.24:9411/api/v2/spans"
	Action      = "publishItem"
)

var (
	fullName = fmt.Sprint(Name, "-", Version, "-", Action)
)

func main() {

	c := pub.Config{Address: Address,
		Name:    Name,
		Version: Version,
		Topic:   operation.InterestOperationV2Topic}

	producer, err := pub.Setup(c)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		addOneWorkItem(producer, c, fmt.Sprint(i))
	}
	producer.Stop()

}

func addOneWorkItem(producer *nsq.Producer, c pub.Config, acc string) zipkin.Ghost {
	start := time.Now()
	item := operation.NewInterestOperationV2(fmt.Sprint(acc))
	s := zipkin.NewSpan(serviceName, fullName)
	item.Ghost = s.ToGhost()

	b, _ := item.ToJsonBytes()
	producer.Publish(c.Topic, b)
	ns := time.Since(start).Nanoseconds()
	ghost := zipkin.LogParentFromSpan(ZipKinUrl, s, ns)
	return ghost

}

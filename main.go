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
)

var (
	fullName              = fmt.Sprint(Name, "-", Version)
	addAllWorkItemsMethod = "addAllWorkItems"
	addOneWorkItemMethod  = "addOneWorkItem"
)

func main() {
	configInterest := pub.Config{Address: Address,
		Name:    Name,
		Version: Version,
		Topic:   operation.InterestOperationV2Topic}

	producerInterest, err := pub.Setup(configInterest)
	if err != nil {
		log.Fatal(err)
	}

	configZipKin := pub.Config{Address: Address,
		Name:    Name,
		Version: Version,
		Topic:   operation.ZipKinOperationV1Topic}

	producerZipKin, err := pub.Setup(configZipKin)
	if err != nil {
		log.Fatal(err)
	}

	addAllWorkItems(producerInterest, configInterest, producerZipKin, configZipKin)
}

func addAllWorkItems(producerInterest *nsq.Producer, configInterest pub.Config, producerZipKin *nsq.Producer, configZipKin pub.Config) {

	start := time.Now()

	for r := 1; r <= 1; r++ {
		for i := 1; i <= 10; i++ {
			addOneWorkItem(producerInterest, configInterest, producerZipKin, configZipKin, fmt.Sprint(i))
		}
	}

	ns := time.Since(start)
	zipkin.LogParent(producerZipKin, configZipKin, fullName, addAllWorkItemsMethod, ns)

	time.Sleep(5 * time.Second)
	producerInterest.Stop()
	producerZipKin.Stop()

}

func addOneWorkItem(producer *nsq.Producer, c pub.Config, zipKinProducer *nsq.Producer, zipkinConfig pub.Config, acc string) zipkin.Ghost {

	start := time.Now()

	item := operation.NewInterestOperationV2(fmt.Sprint(acc))
	s := zipkin.NewSpan(serviceName, addOneWorkItemMethod)
	item.Ghost = s.ToGhost()

	b, _ := item.ToJsonBytes()
	producer.Publish(c.Topic, b)
	d := time.Since(start)
	ghost, _ := zipkin.LogParentSpan(zipKinProducer, zipkinConfig, s, d)
	return ghost

}

package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

var (
	topic      = "interest-calculation-v1"
	nsqAddress = "192.168.88.24:4150"
	name       = "manager"
	version    = "v1"
)

func main() {
	cfg := nsq.NewConfig()
	flag.Var(&nsq.ConfigFlag{cfg}, "producer-opt", "http://godoc.org/github.com/nsqio/go-nsq#Config")
	flag.Parse()
	cfg.UserAgent = fmt.Sprintf("%s-%s", name, version)

	producer, err := nsq.NewProducer(nsqAddress, cfg)
	if err != nil {
		log.Fatalf("failed to create nsq.Producer - %s", err)
	}
	for i := 1; i <= 10; i++ {
		s := fmt.Sprint(i)
		err = producer.Publish(topic, []byte(s))

	}
	producer.Stop()

}

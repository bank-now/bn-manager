package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

var (
	topic      = "interest-calculation"
	nsqAddress = "192.168.88.24:4150"
)

func main() {
	cfg := nsq.NewConfig()
	flag.Var(&nsq.ConfigFlag{cfg}, "producer-opt", "option to passthrough to nsq.Producer (may be given multiple times, http://godoc.org/github.com/nsqio/go-nsq#Config)")

	flag.Parse()

	//termChan := make(chan os.Signal, 1)
	//signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	cfg.UserAgent = fmt.Sprintf("manager-v1")

	// make the producers
	producers := make(map[string]*nsq.Producer)
	producer, err := nsq.NewProducer(nsqAddress, cfg)
	if err != nil {
		log.Fatalf("failed to create nsq.Producer - %s", err)
	}
	producers[nsqAddress] = producer

	for i := 1; i <= 10; i++ {
		s := fmt.Sprint(i)
		err = producer.Publish(topic, []byte(s))

	}

	//select {
	//case <-termChan:
	//}

	for _, producer := range producers {
		producer.Stop()
	}
}

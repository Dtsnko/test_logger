package main

import (
	"context"
	"github.com/webitel/logger/pkg/client"
	"log"
	"time"
)

func main() {
	cl, err := client.NewClient("amqp://webitel:webitel@10.9.8.111:5672", "yehorrr", "10.9.8.111:8500")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = cl.Open()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cl.Close()
	for {
		err = cl.Rabbit().CreateAction(1, 14, 10454, "").One(10, []byte(`{"hello":"hello"}`)).SendContext(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(10 * time.Second)
		err = cl.Rabbit().UpdateAction(
			1,
			14,
			10454,
			"",
		).Many([]int{10, 15, 20}, [][]byte{[]byte(`{"hello":"hello"}`), []byte(`{"hello":"hello"}`), []byte(`{"hello":"hello"}`)}).SendContext(context.Background())
		if err != nil {
			log.Fatal(err.Error())
		}
	}

}

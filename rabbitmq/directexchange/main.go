package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const exchange = "mail"

var (
	binding  = [3]string{"car", "motorcycle", "bicycle"}
	counters = [3]int32{0, 0, 0}
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	if err = ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		panic(err)
	}

	var queues [3]amqp.Queue
	for i := 0; i < 3; i++ {
		queues[i], err = ch.QueueDeclare("", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}

		if err = ch.QueueBind(queues[i].Name, binding[i], exchange, false, nil); err != nil {
			panic(err)
		}
	}

	for i := 0; i < 3; i++ {
		go func(binding string) {
			for {
				if ch.IsClosed() {
					continue
				}
				ctx, cl := context.WithTimeout(context.Background(), 5*time.Second)
				defer cl()
				if err = ch.PublishWithContext(ctx, "mail", binding, false, false, newPublishing(binding)); err != nil {
					panic(err)
				}
				time.Sleep(time.Duration(rand.Intn(10)*100) * time.Millisecond)
			}
		}(binding[i])
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for i := 0; i < 3; i++ {
		go func(i int) {
			c, err := ch.Consume(queues[i].Name, "", true, false, false, false, nil)
			if err != nil {
				return
			}
			for msg := range c {
				atomic.AddInt32(&counters[i], 1)
				log.Println(msg.Exchange, msg.RoutingKey, string(msg.Body))
			}
		}(i)
	}

	<-interrupt
	for i := 0; i < len(counters); i++ {
		log.Println(binding[i], counters[i])
	}
	for _, q := range queues {
		n, err := ch.QueueDelete(q.Name, false, false, false)
		if err != nil {
			log.Println(err)
		}

		log.Println("messages purged for the queue", q.Name, ": ", n)
	}
}

func newPublishing(transport string) amqp.Publishing {
	return amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("this is a message to the " + transport + " mail transport"),
	}
}

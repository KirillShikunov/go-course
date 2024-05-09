package main

import (
	"context"
	"github.com/KirillShikunov/fruit-core/event"
	"github.com/KirillShikunov/fruit-core/queue"
	"github.com/KirillShikunov/fruit-core/rabbitmq"
	"log"
	"os/signal"
	"syscall"
)

type OrangeCounter struct {
	small, medium, large int
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn := rabbitmq.NewConnection("amqp://root:root@localhost:5672/")
	ch, err := conn.GetChannel(ctx)
	if err != nil {
		log.Fatalf("Failed to get channel: %s", err)
	}

	rabbit := rabbitmq.NewChannel(ch)
	queueManager := queue.NewManager(rabbit)

	orangeCounter := &OrangeCounter{}

	err = queueManager.Listen(ctx, "fruit.orange", func(e event.Event) {
		if orangeEvent, ok := e.(*event.OrangeEvent); ok {
			log.Println("Received Orange Event:", orangeEvent.Size)

			switch {
			case orangeEvent.Size < 10:
				orangeCounter.small++
			case orangeEvent.Size < 20:
				orangeCounter.medium++
			default:
				orangeCounter.large++
			}

			log.Printf(
				"Small: %d, Medium: %d, Large: %d\n",
				orangeCounter.small,
				orangeCounter.medium,
				orangeCounter.large,
			)
		}
	})

	<-ctx.Done()

	log.Println("Received shutdown signal, stopping...")
}

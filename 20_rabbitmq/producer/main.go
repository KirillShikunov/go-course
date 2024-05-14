package main

import (
	"context"
	"fmt"
	"github.com/KirillShikunov/fruit-core/event"
	"github.com/KirillShikunov/fruit-core/queue"
	"github.com/KirillShikunov/fruit-core/rabbitmq"
	"log"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

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

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received shutdown signal, stopping...")
			return
		default:
			orangeEvent := &event.OrangeEvent{
				Size: rand.Intn(30),
			}

			if err := queueManager.Post(ctx, orangeEvent); err != nil {
				log.Fatalf("Failed to post event: %s", err)
			}

			fmt.Println("Posted Orange Event:", orangeEvent.Size)

			second := time.Duration(rand.Intn(5))
			time.Sleep(second * time.Second)
		}
	}
}

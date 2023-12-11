package main

import (
	"encoding/json"
	"log"

	"github.com/bondzai/mqsource/mock"
	"github.com/bondzai/mqsource/task"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := declareQueue(ch)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	tasks := mock.GetTasks()
	publishTasks(ch, q.Name, *tasks)
}

func declareQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"tasks",
		false,
		false,
		false,
		false,
		nil,
	)
}

func publishTasks(ch *amqp.Channel, queueName string, tasks []task.Task) {
	for _, t := range tasks {
		taskBytes, err := json.Marshal(t)
		if err != nil {
			log.Printf("Failed to marshal task: %v", err)
			continue
		}

		err = ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        taskBytes,
			},
		)
		if err != nil {
			log.Printf("Failed to publish task: %v", err)
		}

		log.Printf("Task %d published to the queue\n", t.ID)
	}
}

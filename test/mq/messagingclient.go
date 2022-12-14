package mq

import (
	"fmt"
	"log"
	config "test/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	//var wg sync.WaitGroup
	//wg.Add(config.PRODUCERCNT)
	//
	//for routine := 0; routine < config.PRODUCERCNT; routine++ {
	//	go func(routineNum int) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	//conn, err := amqp.Dial(config.RMQADDR)
	q, err := ch.QueueDeclare(
		config.QUEUENAME, //Queue name
		true,             //durable
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	//for i := 0; i < 10; i++ {
	msgBody := fmt.Sprintf("Message__")

	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //Msg set as persistent
			ContentType:  "text/plain",
			Body:         []byte(msgBody),
		})

	log.Printf(" [x] Sent %s", msgBody)
	failOnError(err, "Failed to publish a message")
	//}

	//wg.Done()
	//}(routine)
	//}

	//wg.Wait()

	log.Println("All messages sent!!!!")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}

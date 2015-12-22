package main

import (
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"first.queue", //name string,
		true,          //durable bool,
		false,         //autoDelete bool,
		false,         //exclusive bool,
		false,         //noWait bool,
		nil)           //args amqp.Table)

	msg := amqp.Publishing{
		Body: []byte("my first message"),
	}

	ch.Publish(
		"",     //exchange string,
		q.Name, //key string,
		false,  //mandatory bool,
		false,  //immediate bool,
		msg)    //msg amqp.Publishing)

	msgs, _ := ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  //noLocal bool,
		false,  //noWait bool,
		nil)    //args amqp.Table)

	for m := range msgs {
		println(string(m.Body))
	}
}

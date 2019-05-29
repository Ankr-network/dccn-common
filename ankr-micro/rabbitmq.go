package ankrmicro

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

// RabbitMQHost contains the endpoint of RabbitMQ broker
var RabbitMQHost string

var defaultExchange = "micro"

func failOnError(err error, msg string) {
	if err != nil {
		logStr := fmt.Sprintf("%s: %s", msg, err)
		WriteLog(logStr)
	}
}

func getRabbitMQHost() string {
	config := GetConfig()
	host := config.Rabbitmq
	//logStr := fmt.Sprintf("RabbitMQ hostname : %s", host)
	//WriteLog(logStr)
	return host
}

// Send function send message to RabbitMQ queue
func Send(topic string, e interface{}) {

	conn, err := amqp.Dial(getRabbitMQHost())
	failOnError(err, "Failed to connect to RabbitMQ")

	if conn == nil {
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	if ch == nil {
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		defaultExchange, // name
		"topic",         // type
		false,           // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)

	b, _ := proto.Marshal(e.(proto.Message))

	m := amqp.Publishing{
		Body:    b,
		Headers: amqp.Table{},
	}
	// compatible with go-micro
	m.Headers["Content-Type"] = "application/octet-stream"

	err = ch.Publish(
		defaultExchange, // exchange
		topic,           // routing key
		false,           // mandatory
		false,           // immediate
		m)

	logstr := fmt.Sprintf(" [x] Sent Msg to Topic : %s ", topic)
	WriteLog(logstr)
	failOnError(err, "Failed to publish a message")
}

// Receive function receives messages from RabbitMQ queue,
// support failed reconnect function
func Receive(topic string, handler interface{}) {
	var rabbitCloseError chan *amqp.Error
	WriteLog("RabbitMQ Listening " + topic)

	for {
		var conn *amqp.Connection
		var ch *amqp.Channel
		var err error

		for conn == nil || ch == nil {
			conn, err = amqp.Dial(getRabbitMQHost())
			failOnError(err, "Failed to connect to RabbitMQ")

			if conn == nil {
				WriteLog("sleep 30 seconds then retry connecting")
				time.Sleep(30 * time.Second)
				//conn.Close()
			} else {
				ch, err = conn.Channel()
				failOnError(err, "Failed to open a channel")
				if ch == nil {
					WriteLog("sleep 30 seconds then retry connect channel")
					time.Sleep(30 * time.Second)
				}
			}
		}

		//defer ch.Close()

		err = ch.ExchangeDeclare(
			defaultExchange, // name
			"topic",         // type
			false,           // durable
			false,           // auto-deleted
			false,           // internal
			false,           // no-wait
			nil,             // arguments
		)

		if err != nil {
			log.Panicf(" ExchangeDeclare error  %s \n", err)
		}

		rabbitCloseError = make(chan *amqp.Error)
		conn.NotifyClose(rabbitCloseError)

		q, err := ch.QueueDeclare(
			"",    // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		err = ch.QueueBind(
			q.Name,  // queue name
			topic,   // routing key
			"micro", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		// process handle
		typ := reflect.TypeOf(handler)

		//hdlr := reflect.ValueOf(handler)
		//name := reflect.Indirect(hdlr).Type().Name()

		//	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(0) // only first method will be callback
		methodFunc := method.Func
		//fmt.Printf("method wlll be callback %+v\n", method)
		if method.Type.NumIn() != 2 {
			fmt.Printf("subscribe callback method only have one parameter, real %d \n", method.Type.NumIn())
			return
		}
		reqType := method.Type.In(1)

		//fmt.Printf("reqType %+v\n", reqType)

		go func() {
			for d := range msgs {
				logStr := fmt.Sprintf("Received a Message from Topic : %s ", topic)
				WriteLog(logStr)

				var req reflect.Value
				if reqType.Kind() == reflect.Ptr {
					req = reflect.New(reqType.Elem())
				} else {
					req = reflect.New(reqType)
				}
				newP2 := req.Interface()
				proto.Unmarshal(d.Body, newP2.(proto.Message))

				//fmt.Printf("req %+v \n", newP2)

				var vals []reflect.Value
				handlerValue := reflect.ValueOf(handler)
				vals = append(vals, handlerValue)
				vals = append(vals, reflect.ValueOf(newP2))

				methodFunc.Call(vals)
			}
		}()

		//WriteLog(" [*] Waiting for messages. To exit press CTRL+C")
		msg := <-rabbitCloseError
		//conn.Close()
		logStr := fmt.Sprintf("receive rabbitMQ close messages, error : %s ", msg.Reason)
		WriteLog(logStr)
	}

}

package utils

import (
	"fmt"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var RabbitCh *amqp.Channel

func ConnectToRabbitMQ() (err error) {
	config := Config.RabbitMQConfig
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Account, config.Password, config.Host, config.Port)
	// fmt.Println(rabbitMQURL, rabbitMQURL == consts.RabbitMQURL)
	if Conn, err = amqp.Dial(rabbitMQURL); err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	if RabbitCh, err = Conn.Channel(); err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	return nil
}

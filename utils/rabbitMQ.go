package utils

import (
	"usermanagersystem/consts"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var RabbitCh *amqp.Channel

func ConnectToRabbitMQ() (err error) {
	if Conn, err = amqp.Dial(consts.RabbitMQURL); err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	if RabbitCh, err = Conn.Channel(); err != nil {
		return ErrWrapOrWithMessage(true, err)
	}
	return nil
}

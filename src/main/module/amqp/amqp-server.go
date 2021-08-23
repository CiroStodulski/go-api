package amqp_server

import (
	"go-api/src/main/container"
	consumer "go-api/src/presentation/amqp/consumers"
)

type AmqpServer interface {
	New(c *container.Container) AmqpServer
	Start()
	StartConsumers()
	NeedToReconnect(err error, msg string)
	LoadConsumers(c *container.Container) []consumer.Comsumer
	LoadProducers(c *container.Container)
}

package consumer

import (
	"github.com/nats-io/nats.go"
	"log"
)

const (
	URL            = "jupyterhub.brique.kr:4222"
	subSubjectName = "TEST.NATS"

	User = "myuser"
	Pass = "mypass"
)

type Consumer struct {
	nc  *nats.Conn
	js  nats.JetStreamContext
	sub *nats.Subscription
}

func NewConsumer() (*Consumer, error) {
	nc, err := nats.Connect(URL, nats.UserInfo(User, Pass))
	if err != nil {
		return nil, err
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	sub, err := js.PullSubscribe(subSubjectName, "hello")
	if err != nil {
		return nil, err
	}

	return &Consumer{nc: nc, js: js, sub: sub}, nil
}

func (c *Consumer) Close() {
	c.nc.Close()
}

func (c *Consumer) ClaimMessage() error {
	msgs, err := c.sub.Fetch(1)
	if err == nats.ErrTimeout {

	} else if err != nil {
		return err
	}

	for _, msg := range msgs {
		if err := msg.Ack(); err != nil {
			return err
		}
		log.Printf("Data: %s\n", string(msg.Data))
	}
	return nil
}

func (c *Consumer) Unsubscribe() {
	if err := c.sub.Unsubscribe(); err != nil {
		log.Println(err)
	}
}

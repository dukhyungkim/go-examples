package producer

import (
	"github.com/nats-io/nats.go"
	"log"
)

const (
	URL         = "jupyterhub.brique.kr:4222"
	subjectName = "TEST.NATS"

	User = "myuser"
	Pass = "mypass"
)

type Producer struct {
	nc *nats.Conn
	js nats.JetStreamContext
}

func NewProducer() (*Producer, error) {
	nc, err := nats.Connect(URL, nats.UserInfo(User, Pass))
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &Producer{nc: nc, js: js}, nil
}

func (p *Producer) Close() {
	p.nc.Close()
}

func (p *Producer) CreateStream(streamName string) error {
	streamInfo, err := p.js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	streamSubjects := streamName + ".*"
	if streamInfo == nil {
		log.Printf("create stream %q and subjects %q", streamName, streamSubjects)
		streamConfig := nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		}
		if _, err := p.js.AddStream(&streamConfig); err != nil {
			return err
		}
	}
	return nil
}

func (p *Producer) SendMessage() error {
	pub, err := p.js.Publish(subjectName, []byte("hello"))
	if err != nil {
		return err
	}

	log.Printf("%+v\n", pub)
	return nil
}

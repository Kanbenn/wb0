package pubsub

import (
	"log"

	"github.com/Kanbenn/mywbgonats/internal/config"
	"github.com/nats-io/stan.go"
)

type NatsCon struct {
	cfg     config.Config
	con     stan.Conn
	sub     stan.Subscription
	next    processer
	subject string
	durable string
}

type processer interface {
	ProcessNatsMessage([]byte)
}

func Connect(cluster, addr, clientID string) *NatsCon {
	sc, err := stan.Connect(cluster, clientID, stan.NatsURL(addr))
	if err != nil {
		log.Fatal("error at connecting to nats", err)
	}
	log.Println("connected to nats-stream as", clientID)
	return &NatsCon{con: sc}
}

func (n *NatsCon) RegisterMessageProcessor(next processer) {
	n.next = next
}

func (n *NatsCon) StartListeningForNewMessages() {
	// опция DurableName позволяет получить пропущенные сообщения
	// при пере-подключении к серверу nats.
	ss, err := n.con.Subscribe(
		n.cfg.NatsSubject,
		n.recieveNatsMsg,
		stan.DurableName(n.durable))
	if err != nil {
		log.Fatal("error at subscribing to nats", err, n.cfg)
	}
	n.sub = ss
}

func (n *NatsCon) Publish(data []byte) {
	log.Println("publishing new msg to nats")
	n.con.Publish(n.cfg.NatsSubject, data)
}

func (n *NatsCon) Close() {
	if n.sub != nil {
		log.Println("closing nats subscription", n.sub.Close())
	}
	log.Println("closing nats connection", n.con.Close())
}

func (n *NatsCon) recieveNatsMsg(m *stan.Msg) {
	log.Println("got new msg from nats", m.Size(), m.Timestamp)
	n.next.ProcessNatsMessage(m.Data)
}

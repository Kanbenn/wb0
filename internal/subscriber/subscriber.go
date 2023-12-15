package subscriber

import (
	"log"

	"github.com/Kanbenn/mywbgonats/internal/config"
	"github.com/nats-io/stan.go"
)

type NatsCon struct {
	cfg  config.Config
	con  stan.Conn
	sub  stan.Subscription
	next processer
}

type processer interface {
	ProcessNatsMessage([]byte)
}

func New(cfg config.Config, next processer) *NatsCon {
	sc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL(cfg.Nats))
	if err != nil {
		log.Fatal("error at connecting to nats", err)
	}

	n := NatsCon{cfg: cfg, con: sc, next: next}

	return &n
}

func (n *NatsCon) SubscribeOnSubject() {
	// опция DurableName позволяет получить пропущенные сообщения
	// при пере-подключении к серверу nats.
	ss, err := n.con.Subscribe(
		n.cfg.NatsSubject,
		n.recieveNatsMsg,
		stan.DurableName(n.cfg.NatsDurable))
	if err != nil {
		log.Fatal("error at subscribing to nats", err, n.cfg.NatsSubject, n.cfg.NatsDurable)
	}
	n.sub = ss
}

func (n *NatsCon) Close() {
	n.sub.Close()
	n.con.Close()
}

func (n *NatsCon) recieveNatsMsg(m *stan.Msg) {
	log.Println("handleMsg: got new msg from nats", m.Size(), m.Timestamp)
	n.next.ProcessNatsMessage(m.Data)
}

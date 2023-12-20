package config

import (
	"flag"
)

type Config struct {
	Addr        string
	Nats        string
	PgConnStr   string
	NatsSubject string
	NatsDurable string
	NatsCluster string
}

func New() Config {
	c := Config{}
	c.NatsSubject = "wb-orders"
	c.NatsDurable = "my-durable"
	c.NatsCluster = "test-cluster"
	return c
}
func (c *Config) ParseFlags() {
	flag.StringVar(&c.Addr, "a", "localhost:8080", "web-server run address")
	flag.StringVar(&c.Nats, "n", "nats://localhost:4222", "nats-stream address to subcribe to")
	flag.StringVar(&c.PgConnStr, "p", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "Postgres connection string")
	flag.Parse()
}

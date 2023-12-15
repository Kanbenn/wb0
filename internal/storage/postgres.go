package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Kanbenn/mywbgonats/internal/config"
	"github.com/Kanbenn/mywbgonats/internal/models"
)

const createTable = `
	CREATE TABLE IF NOT EXISTS orders (
		id 		   SERIAL PRIMARY KEY,
		order_uid  VARCHAR(100) UNIQUE,
		order_data JSONB);
	CREATE INDEX IF NOT EXISTS oid ON orders(order_uid);`

type Pg struct {
	Sqlx *sqlx.DB
	Cfg  config.Config
}

func NewPostgres(cfg config.Config) *Pg {
	conn, err := sqlx.Open("postgres", cfg.PgConnStr)
	if err != nil {
		log.Fatal("error at connecting to Postgres:", cfg.PgConnStr, err)
	}

	pg := Pg{conn, cfg}

	if _, err := pg.Sqlx.Exec(createTable); err != nil {
		log.Println("error at creating db-tables:", pg.Cfg.PgConnStr, conn)
		log.Fatal(err)
	}
	return &pg
}

func (pg *Pg) Close() error {
	return pg.Sqlx.Close()
}

func (pg *Pg) InsertOrder(o models.Order) {
	q := "INSERT INTO orders (order_uid, order_data) VALUES(:order_uid, :order_data)"
	_, err := pg.Sqlx.NamedExec(q, o)
	if err != nil {
		log.Println("pg.InsertOrder error:", err)
	}
}

func (pg *Pg) SelectAllOrders() (orders []models.Order) {
	q := "SELECT order_uid, order_data FROM orders"
	err := pg.Sqlx.Select(&orders, q)
	if err != nil {
		log.Println("pg.SelectAllOrders error:", err)
	}
	return orders
}

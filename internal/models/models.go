package models

type Order struct {
	ID   string `db:"order_uid"  json:"order_uid"`
	Data []byte `db:"order_data" json:"-"`
}

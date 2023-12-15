package app

import (
	"encoding/json"
	"log"

	"github.com/Kanbenn/mywbgonats/internal/models"
	"github.com/Kanbenn/mywbgonats/internal/storage"
)

type App struct {
	// Cfg config.Config
	// s   storer
	Ch *storage.Cache
	Pg *storage.Pg
}

// func (app *App) processNewOrder(data []byte) {

// }

func New(ch *storage.Cache, pg *storage.Pg) *App {
	return &App{ch, pg}
}

func (app *App) RestoreCacheDataFromPg() {
	orders := app.Pg.SelectAllOrders()
	app.Ch.AddBatch(orders)
}

func (app *App) ProcessNatsMessage(data []byte) {
	o := models.Order{}
	if err := json.Unmarshal(data, &o); err != nil {
		log.Println("ProcessNatsMessage: error at unmarshalling the data", err)
		return
	}
	if len(o.ID) < 1 {
		log.Println("ProcessNatsMessage error: order_uid tag was not found in the input data")
		return
	}
	o.Data = data

	app.Ch.Add(o.ID, o.Data)
	log.Println("Cache contents after adding new record:", app.Ch)

	app.Pg.InsertOrder(o)
}

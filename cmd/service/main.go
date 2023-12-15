package main

import (
	"github.com/Kanbenn/mywbgonats/internal/app"
	"github.com/Kanbenn/mywbgonats/internal/config"
	"github.com/Kanbenn/mywbgonats/internal/storage"
	"github.com/Kanbenn/mywbgonats/internal/subscriber"
	"github.com/Kanbenn/mywbgonats/internal/webserver"
)

func main() {
	cfg := config.NewFromFlags()

	ch := storage.NewCache()
	pg := storage.NewPostgres(cfg)
	defer pg.Close()

	app := app.New(ch, pg)
	app.RestoreCacheDataFromPg()

	sub := subscriber.New(cfg, app)
	sub.SubscribeOnSubject()
	defer sub.Close()

	srv := webserver.New(cfg, app)
	go srv.ShutdownOnSignal()
	srv.Launch()

}

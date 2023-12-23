package main

import (
	"github.com/Kanbenn/mywbgonats/internal/app"
	"github.com/Kanbenn/mywbgonats/internal/config"
	"github.com/Kanbenn/mywbgonats/internal/pubsub"
	"github.com/Kanbenn/mywbgonats/internal/storage"
	"github.com/Kanbenn/mywbgonats/internal/webserver"
)

func main() {
	cfg := config.New()
	cfg.ParseFlags()

	ch := storage.NewCache()
	pg := storage.NewPostgres(cfg)
	defer pg.Close()

	app := app.New(ch, pg)
	app.RestoreCacheDataFromPg()

	sub := pubsub.New(cfg, "subscriber")
	sub.SubscribeOnSubject(app)
	defer sub.Close()

	srv := webserver.New(cfg, ch)
	go srv.ShutdownOnSignal()
	srv.Launch()

}

package server

import (
	"flag"
	"github.com/penguinn/penguin/component/config"
	"github.com/penguinn/penguin/component/db"
	"github.com/penguinn/penguin/component/log"
	"github.com/penguinn/penguin/component/router"
)

func init() {
	flag.Parse()

	Use(config.ConfigComponent{})
	Use(log.LogComponent{})
	Use(router.RouterComponent{})
	Use(db.DBComponent{})
}

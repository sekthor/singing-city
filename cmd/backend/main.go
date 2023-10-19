package main

import (
	"github.com/sekthor/songbird-backend/pkg/api"
	"github.com/sekthor/songbird-backend/pkg/config"
)

func main() {
	conf := config.LoadConfig()
	api, err := api.NewApi(conf)

	if err != nil {
		// TODO: graceful shutdown
	}

	router := api.Router()
	err = router.Run(conf.Server.Host + ":" + conf.Server.Port)

	if err != nil {
		// TODO: graceful shutdown
	}
}

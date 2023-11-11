package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/api"
	"github.com/sekthor/songbird-backend/pkg/config"
)

func main() {
	conf := config.LoadConfig()
	api, err := api.NewApi(conf)

	if err != nil {
		log.Fatal().Err(err).Msg("could not create api")
	}

	router := api.Router()
	err = router.Run(conf.Server.Host + ":" + conf.Server.Port)

	if err != nil {
		log.Fatal().Err(err).Msg("could start router")
	}
}

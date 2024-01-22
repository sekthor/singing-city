package main

import (
	"github.com/rs/zerolog/log"
	"github.com/sekthor/singing-city/pkg/api"
	"github.com/sekthor/singing-city/pkg/config"
)

func main() {

	log.Info().Msg("loading configuration")
	conf := config.LoadConfig()

	log.Info().Msg("creating api")
	api, err := api.NewApi(conf)

	if err != nil {
		log.Fatal().Err(err).Msg("could not create api")
	}

	log.Info().Msg("creating router")
	router := api.Router()

	log.Info().Msg("starting router running at '" + conf.Server.Host + ":" + conf.Server.Port + "'")
	err = router.Run(conf.Server.Host + ":" + conf.Server.Port)

	if err != nil {
		log.Fatal().Err(err).Msg("could start router")
	}
}

package pkg

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	STAGE          string
	PORT           string
	OPENAI_API_KEY string `validate:"required"`
}

func LoadEnvs() (config EnvVars, error error) {
	STAGE := os.Getenv("GO_ENV")

	if STAGE == "" {
		STAGE = "dev"
	}

	if STAGE != "prod" {
		err := godotenv.Load(".env." + STAGE)
		if err != nil {
			log.Fatal().Msg("fail load env file, set required envs in prod")
			return config, err
		}
	}

	config.PORT = os.Getenv("PORT")
	config.OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

	log.Info().Any("config", config)

	// Validate config
	errs := Valtor.Validate(config)
	if errs != nil {
		log.Warn().Msg("error validating config:" + errs.Error())
		return config, errs
	}
	return
}

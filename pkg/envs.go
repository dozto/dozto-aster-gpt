package pkg

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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

	// Load config from file for non prod env
	if STAGE != "prod" {
		viper.AddConfigPath("./")
		viper.SetConfigName(".env." + STAGE)

		// Read config file
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				log.Fatal().Msg("config file not found.")
				// log.Error("!! Config file not found.")
			} else {
				log.Fatal().Msg("error reading config file:" + err.Error())
			}
			return config, err
		}
	}

	// Automatically read env variables
	viper.AutomaticEnv()

	// Unmarshal config
	viper.Unmarshal(&config)

	log.Info().Msg("env Config loaded success.")
	// // Validate config
	errs := Valtor.Validate(config)
	if errs != nil {
		log.Warn().Msg("error validating config:" + errs.Error())
		return config, errs
	}

	return
}

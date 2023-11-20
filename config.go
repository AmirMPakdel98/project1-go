package main

import (
	"c-vod/utils/helper"
	"c-vod/utils/types"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*types.AppConfig, error) {

	ext_path := helper.GetCurrentDir()

	env_file_name := ".env.prod"

	if os.Getenv("ENV") == "dev" {
		env_file_name = ".env.dev"
	}

	env_path := ext_path + "/" + env_file_name

	log.Println("App running dir : " + ext_path)

	err := godotenv.Load(env_path)

	if err != nil {
		return nil, errors.New("error in loading .env file")
	}

	return &types.AppConfig{
		Env:               os.Getenv("ENV"),
		Log_enabled:       os.Getenv("LOG_ENABLED"),
		App_port:          os.Getenv("APP_PORT"),
		App_domain:        os.Getenv("APP_DOMAIN"),
		App_api_prefix_v1: os.Getenv("APP_API_PREFIX_V1"),

		Database_host:     os.Getenv("DATABASE_HOST"),
		Database_port:     os.Getenv("DATABASE_PORT"),
		Database_name:     os.Getenv("DATABASE_NAME"),
		Database_username: os.Getenv("DATABASE_USERNAME"),
		Database_password: os.Getenv("DATABASE_PASSWORD"),
	}, nil
}

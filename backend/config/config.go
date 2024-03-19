package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *Config

type DbConfig struct {
	Driver string
	URL    string
}
type Config struct {
	HttpPort   string
	SwaggerUrl string
	Db         DbConfig
}

func GetConfig() Config {
	if cfg == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println(err.Error())
		}
		dbUrl := os.Getenv("DB_MAIN_URL")
		httpPort := os.Getenv("HTTP_PORT")
		swgUrl := os.Getenv("SWAGGER_URL")
		if dbUrl == "" {
			dbUrl = "user=root password=root dbname=toggl host=localhost port=6000 sslmode=disable"
		}

		if httpPort == "" {
			httpPort = "8085"
		}

		if swgUrl == "" {
			swgUrl = fmt.Sprintf("http://localhost:%v", httpPort)
		}
		cfg = &Config{
			HttpPort:   httpPort,
			SwaggerUrl: swgUrl,
			Db: DbConfig{
				Driver: "postgres",
				URL:    dbUrl,
			},
		}
	}

	return *cfg
}

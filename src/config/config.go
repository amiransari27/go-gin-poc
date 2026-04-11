package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"

	"github.com/spf13/viper"
)

type Configurations struct {
	Mongo       MongoConfigurations
	Mysql       MysqlConfigurations
	LogDir      string
	Port        string
	RateLimiter RateLimiterConfigurations
}

type MongoConfigurations struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
}

type MysqlConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

type RateLimiterConfigurations struct {
	Enabled bool
	Rate    rate.Limit
	Burst   int
}

var configuration = &Configurations{}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found")
	}

	env := strings.ToUpper(os.Getenv("ENV"))
	fmt.Println("Env : ", env)

	if env == "PROD" {
		viper.SetConfigName("config.prod.yml")
	} else if env == "STAGE" {
		viper.SetConfigName("config.stage.yml")
	} else {
		viper.SetConfigName("config.local.yml")
	}

	viper.AddConfigPath(".")
	// viper.AutomaticEnv() // Enable VIPER to read Environment Variables
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	// Set undefined variables
	// viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(configuration)
	if err != nil {
		log.Fatal("Unable to decode into struct", err)
	}

	// Initialize MongoConfigurations with .env data
	configuration.Mongo = MongoConfigurations{
		DBHost:     os.Getenv("MONGO_HOST"),
		DBName:     os.Getenv("MONGO_DBNAME"),
		DBUser:     os.Getenv("MONGO_USERNAME"),
		DBPassword: os.Getenv("MONGO_PASSWORD"),
	}

}

func GetConfig() *Configurations {

	return configuration
}

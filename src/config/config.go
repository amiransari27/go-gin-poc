package config

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Configurations struct {
	Mongo  MongoConfigurations
	Mysql  MysqlConfigurations
	LogDir string
	Port   string
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

var configuration = &Configurations{}

func init() {
	env := strings.ToUpper(os.Getenv("env"))
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
}

func GetConfig() *Configurations {

	return configuration
}

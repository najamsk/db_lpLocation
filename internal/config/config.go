package config

import (
	"fmt"
	"log"

	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Constants do we really need to export them? maybe for marshal
type constants struct {
	MoftakDevEmails []string
	Environment     string
	PORT            string
	BasicAuth       struct {
		Username string
		Password string
	}
	JWT struct {
		JWTPASS string
	}
	CockroachDB struct {
		Host            string
		Port            string
		Database        string
		DefaultDatabase string
		User            string
		SSL             bool
	}
	SMTP struct {
		Host      string
		Port      int
		User      string
		Password  string
		EmailFrom string
	}
	Contact struct {
		Phone      string
		Dev        string
		Email      string
		Web        string
		Phone2     string
		WebDisplay string
	}
}

//Config will hold our constants as items
type Config struct {
	// Database *mgo.Database
	Items constants
	// Database *mgo.Database
}

// New func: NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Items = constants
	if err != nil {
		return &config, err
	}
	// dbSession, err := mgo.Dial(config.Constants.Mongo.URL)
	// if err != nil {
	// 	return &config, err
	// }
	// config.Database = dbSession.DB(config.Constants.Mongo.DBName)
	return &config, err
}

// maybe pass config location using command line argument using flag package
// https://stackoverflow.com/questions/35419263/using-a-configuration-file-with-a-compiled-go-program
func initViper() (constants, error) {
	fmt.Println("env mode is :")
	fmt.Println(os.Getenv("mode"))
	if os.Getenv("mode") == "staging" {
		viper.SetConfigName("settings.stagconfig") // Configuration fileName without the .TOML or .YAML extension
	} else if os.Getenv("mode") == "live" {
		viper.SetConfigName("settings.config.live")
	} else if os.Getenv("mode") == "dev" {
		viper.SetConfigName("settings.config.dev")
	} else {
		viper.SetConfigName("settings.config") // Configuration fileName without the .TOML or .YAML extension
	}

	viper.AddConfigPath("./internal/config") // Search the root directory for the configuration file
	viper.AddConfigPath("./")                // Search the root directory for the configuration file
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		return constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants constants
	err = viper.Unmarshal(&constants)

	// here get values from secure .env file and override constants
	jwtpass := os.Getenv("JWTPASS")
	if len(jwtpass) > 0 {
		constants.JWT.JWTPASS = os.Getenv("JWTPASS")
	}
	return constants, err
}

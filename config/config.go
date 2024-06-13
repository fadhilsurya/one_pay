package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JwtConfig
	DB       *gorm.DB
}

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

type ServerConfig struct {
	Port int
}

type JwtConfig struct {
	Secret string
}

var AppConfig Config

func InitCofig() {

	// examples of viper if we want to use .env file
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()
	// fmt.Println(viper.Get("PORT"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	databaseConfig := AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name)

	// Initialize GORM with MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	AppConfig.DB = db

	fmt.Println("Database connection successful!")

}

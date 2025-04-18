package user

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

func loadConfig() DBConfig {
	cfg := DBConfig{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
	}

	if cfg.Host == "" || cfg.User == "" || cfg.Pass == "" || cfg.Name == "" || cfg.Port == "" {
		log.Fatal("Missing one or more required DB environment variables")
	}

	return cfg
}

func InitDB() {
	_ = godotenv.Load(".env")

	config := loadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Pass, config.Name, config.Port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
}

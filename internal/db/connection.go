package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/thechauhanabhay/blabberit/internal/message"
	"github.com/thechauhanabhay/blabberit/internal/user"
)

var (
	db   *gorm.DB
	once sync.Once
)

type DBConfig struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

func loadConfig() DBConfig {
	_ = godotenv.Load(".env")

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

func Init() *gorm.DB {
	once.Do(func() {
		cfg := loadConfig()
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
			cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Port,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err := db.AutoMigrate(
			&user.User{},
			&message.Message{},
		); err != nil {
			log.Fatalf("Database migration failed: %v", err)
		}

		user.SetDB(db)
		message.SetDB(db)

		log.Println("Database initialised and migrated")
	})
	return db
}

func Get() *gorm.DB {
	if db == nil {
		return Init()
	}
	return db
}

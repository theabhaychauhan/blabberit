package testutils

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupTestDB(dsn string, models ...interface{}) *gorm.DB {
	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		panic("could not connect to test database: " + err.Error())
	}

	DB.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
	DB.AutoMigrate(models...)
	return DB
}

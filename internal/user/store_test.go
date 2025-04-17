package user

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestInitDB(t *testing.T) {
	err := godotenv.Load("../../.env")

	if err != nil {
		t.Fatal("Failed to load .env:", err)
	}

	InitDB()

	if DB == nil {
		t.Fatal("Expected DB to be Initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		t.Fatal("Failed to get raw DB Handler")
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("DB connection failed: %v", err)
	}
}

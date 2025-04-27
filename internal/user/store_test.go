// internal/user/store_test.go
package user_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/thechauhanabhay/blabberit/internal/db"
	"github.com/thechauhanabhay/blabberit/internal/user"
)

func TestInitDB(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatal("failed to load .env:", err)
	}

	db.Init()

	if user.DB == nil {
		t.Fatal("expected user.DB to be injected by db.Init()")
	}

	sqlDB, err := user.DB.DB()
	if err != nil {
		t.Fatal("failed to get raw DB handle:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		t.Fatalf("DB ping failed: %v", err)
	}
}

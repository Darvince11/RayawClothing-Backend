package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	//load envs
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Fatal(err)
	}

	dbUrl := os.Getenv("DATABASE_URL")

	//connect to db
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		t.Fatalf("expected no error, got:%v", err)
	}

	return db
}

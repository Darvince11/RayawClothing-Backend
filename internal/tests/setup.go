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
	if dbUrl == "" {
		t.Fatal("DATABASE_URL is not set in the environment")
	}

	//connect to db
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		t.Fatalf("expected no error connecting to database, got:%v", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	t.Log("Database connection has been setup successfully")
	return db
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rayaw-api/internal/config"
	"rayaw-api/internal/routes"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	//load config
	cfg := config.Init()
	//connect to database
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	//start api server
	port := ":" + cfg.Port
	handler := http.TimeoutHandler(routes.ServerMux(cfg, db), 20*time.Second, "Request timeout")

	server := http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	fmt.Printf("Server is listening on port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

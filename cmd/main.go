package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rayaw-api/internal/routes"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	//connect to database
	connstr := "postgresql://rayaw_database_puvg_user:ctayXJGV4yPMe7GU8r0PXiUF1BXCUs87@dpg-d559hjqli9vc73ceuhqg-a.oregon-postgres.render.com/rayaw_database_puvg"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	//start api server
	port := ":8080"
	handler := http.TimeoutHandler(routes.ServerMux(), 20*time.Second, "Request timeout")

	server := http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	fmt.Printf("Server is listening on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

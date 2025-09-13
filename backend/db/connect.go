package db

import (
    "database/sql"
    "log"
    "os"
    "time"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dbURL := os.Getenv("DB_URL")
    var db *sql.DB
    var err error

    for i := 0; i < 10; i++ {
        db, err = sql.Open("postgres", dbURL)
        if err == nil {
            err = db.Ping()
            if err == nil {
                DB = db 
                log.Println("✅ Connected to DB")
                return
            }
        }
        log.Printf("⏳ Attempt %d: Waiting for DB to be ready...", i+1)
        time.Sleep(2 * time.Second)
    }

    log.Fatalf("❌ Failed to connect to DB: %v", err)
}

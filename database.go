package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd:   os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "shopper_db",
        AllowNativePasswords: true,
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Database Connected")
}

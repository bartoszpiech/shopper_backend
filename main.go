package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
    id      int64
    name    string
    email   string
}

type Item struct {
    id          int64
    owner_id    int64
    name        string
    size        string
    link        string
}

func main() {
    ConnectDB()
    //AddUser(User{ name: "Bartek", email: "shaggysamp@interia.pl" })
    users, err := GetUsers()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Users found: %v\n", users)
}

func ConnectDB() {
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

func GetUsers() ([]User, error){
    var users []User
    rows, err := db.Query("SELECT * FROM USER")
    if err != nil {
        return nil, fmt.Errorf("GetUsers: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.id, &user.name, &user.email); err != nil {
            return nil, fmt.Errorf("GetUsers: %v", err)
        }
        users = append(users, user)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetUsers: %v", err)
    }
    return users, nil
}

func AddUser(user User) (int64, error) {
    result, err := db.Exec("INSERT INTO USER (name, email) VALUES (?, ?)", user.name, user.email)
    if err != nil {
        return 0, fmt.Errorf("AddUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddUser: %v", err)
    }
    return id, nil
}

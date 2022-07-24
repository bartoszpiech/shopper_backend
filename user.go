package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    Id      int64       `json:"id"`
    Name    string      `json:"name"`
    Email   string      `json:"email"`
}

func GetUsers(c *gin.Context) {
    var users []User
    rows, err := db.Query("SELECT * FROM USER")
    if err != nil {
        fmt.Errorf("GetUsers: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
            fmt.Errorf("GetUsers: %v", err)
        }
        users = append(users, user)
    }
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.IndentedJSON(http.StatusOK, users)
}

func AddUser(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    var newUser User
    if err := c.BindJSON(&newUser); err != nil {
        return
    }
    result, err := db.Exec("INSERT INTO USER (name, email) VALUES (?, ?)", newUser.Name, newUser.Email)
    if err != nil {
        fmt.Errorf("AddUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Errorf("AddUser: %v", err)
    }
    c.IndentedJSON(http.StatusCreated, gin.H{ "id": id})
}

func GetUserByID(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    var user User
    id := c.Param("id")
    row := db.QueryRow("SELECT * FROM USER WHERE id = ?", id)
    if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
        if err == sql.ErrNoRows {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
            fmt.Errorf("GetUserByID %d: no such user", id)
            return
        }
        fmt.Errorf("GetUserByID %d: %v", id, err)
    }
    c.IndentedJSON(http.StatusOK, user)
}

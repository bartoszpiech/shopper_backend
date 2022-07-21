package main

import (
    "fmt"
)

type User struct {
    id      int64
    name    string
    email   string
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

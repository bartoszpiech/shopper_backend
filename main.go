package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    Connect()

    router := gin.Default()

    router.StaticFS("/static", http.Dir("static"))
    router.GET("/users", GetUsers)

    router.Run("localhost:8080")

    //AddUser(User{ name: "Bartek", email: "shaggysamp@interia.pl" })
    users, err := GetUsers()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Users found: %v\n", users)

    //AddItem(Item{ name: "Buty", size: "42", link: "facebook.com" }, users[2])
    //items, err := GetUserItems(users[2])
    items, err := GetUserItemsByID(1)
    //items, err := GetItems()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Items found: %v\n", items)
}

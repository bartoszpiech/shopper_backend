package main

import (
    "fmt"
    "log"
)

func main() {
    ConnectDB()
    //AddUser(User{ name: "Bartek", email: "shaggysamp@interia.pl" })
    users, err := GetUsers()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Users found: %v\n", users)
    fmt.Printf("\n\n%v\n\n", users[2])

    AddItem(Item{ name: "Buty", size: "42", link: "facebook.com" }, users[2])
    items, err := GetItems()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Items found: %v\n", items)
}


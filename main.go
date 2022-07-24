package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Auth struct {
    Password string `json:"pass"`
}

func main() {
    Connect()

    router := gin.Default()

    router.StaticFS("/static", http.Dir("static"))
    router.GET("/users", GetUsers)
    router.GET("/user/:id", GetUserByID)
    router.POST("/newuser", AddUser)
    router.GET("/user/:id/items", GetUserItems)

    router.GET("/items", GetItems)
    router.POST("/newitem", AddItem)

    router.POST("/auth", Authentication)

    router.Run("localhost:8080")
}

func Authentication(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    var newAuth Auth
    if err := c.BindJSON(&newAuth); err != nil {
        return
    }
    if newAuth.Password == "siema" {
        cookie, err := c.Cookie("gin_cookie")
        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
            fmt.Printf("Cookie value: %s \n", cookie)
            c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong password"})
}

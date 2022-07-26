package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "io"
    "os/exec"
    "github.com/gin-gonic/gin"
    "github.com/mileusna/crontab"
)

type CheckedItem struct {
    AvailableSizes  []string    `json:"available sizes"`
    Description     string      `json:"description"`
    Link            string      `json:"link"`
    Name            string      `json:"name"`
    OldPrice        float64     `json:"old price"`
    Price           float64     `json:"price"`
    Sizes           []string    `json:"sizes"`
}

func CheckSize(item CheckedItem, size string) bool {
    for _, v := range item.AvailableSizes {
        if v == size {
            return true
        }
    }
    return false
}
func CheckItems() {
    users, err := GetAllUsers()
    if err != nil {
        fmt.Println(err)
    }
    for _, i := range users {
        userItems, err := GetAllUserItems(i.Id)
        if err != nil {
            fmt.Println(err)
        }
        var foundItems []Item
        for _, j := range userItems {
            resp, err := http.Get("http://localhost:5000/item/" + j.Link)
            if err != nil {
                fmt.Println(err)
            }
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                fmt.Println(err)
            }
            var checkedItem CheckedItem
            err = json.Unmarshal([]byte(body), &checkedItem)
            if err != nil {
                fmt.Println(err)
            }
            if CheckSize(checkedItem, j.Size) {
                foundItems = append(foundItems, j)
            }
        }
        var message string = "Witaj " + i.Name + ", poniższe przedmioty są ponownie dostępne:"
        message += "\n"
        for _, k := range foundItems {
            message += k.Name
            message += "\t"
            message += k.Size
            message += "\t"
            message += k.Link
            message += "\n"
        }
        fmt.Println("sent email to " + i.Name)
        cmd := exec.Command("neomutt", "-s", "Zara shopper -- twój prywatny asystent zakupowy", i.Email)
        stdin, err := cmd.StdinPipe()
        if err != nil {
            fmt.Println(err)
        }
        go func() {
            defer stdin.Close()
            io.WriteString(stdin, message)
        }()

        _, err = cmd.CombinedOutput()
        if err != nil {
            fmt.Println(err)
        }
    }
}

func main() {
    Connect()

    ctab := crontab.New()
    ctab.MustAddJob("*/5 * * * *", CheckItems)

    CheckItems()
    router := gin.Default()

    router.StaticFS("/static", http.Dir("static"))
    router.GET("/users", GetUsers)
    router.GET("/users/:id", GetUserByID)
    router.POST("/newuser", AddUser)
    router.GET("/users/:id/items", GetUserItems)
    router.GET("/deleteuser/:id", DeleteUser)

    router.GET("/items", GetItems)
    router.POST("/newitem", AddItem)
    router.GET("/deleteitem/:id", DeleteItem)

    //router.POST("/auth", Authentication)

    router.Run("localhost:8080")
}

/*
type Auth struct {
    Password string `json:"pass"`
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
            c.SetCookie("gin_cookie", "test", 3600, "/auth", "localhost", false, true)
            fmt.Printf("Cookie value: %s \n", cookie)
            c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong password"})
}
*/

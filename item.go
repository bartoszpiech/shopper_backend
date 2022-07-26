package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Item struct {
    Id          int64   `json:"id"`
    Owner_id    int64   `json:"owner_id"`
    Name        string  `json:"name"`
    Size        string  `json:"size"`
    Link        string  `json:"link"`
}

func GetAllItems() ([]Item, error){
    var items []Item
    rows, err := db.Query("SELECT * FROM ITEM")
    if err != nil {
        return nil, fmt.Errorf("GetAllItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.Id, &item.Owner_id, &item.Name, &item.Size, &item.Link); err != nil {
            return nil, fmt.Errorf("GetAllItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetAllItems: %v", err)
    }
    return items, nil
}

func GetItems(c *gin.Context) {
    var items []Item
    rows, err := db.Query("SELECT * FROM ITEM")
    if err != nil {
        fmt.Errorf("GetAllItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.Id, &item.Owner_id, &item.Name, &item.Size, &item.Link); err != nil {
            fmt.Errorf("GetItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        fmt.Errorf("GetItems: %v", err)
    }
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.IndentedJSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    var newItem Item
    if err := c.BindJSON(&newItem); err != nil {
        return
    }
    result, err := db.Exec("INSERT INTO ITEM (owner_id, name, size, link) VALUES (?, ?, ?, ?)", newItem.Owner_id, newItem.Name, newItem.Size, newItem.Link)
    if err != nil {
        fmt.Errorf("AddItem: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Errorf("AddItem: %v", err)
    }
    c.IndentedJSON(http.StatusCreated, gin.H{ "id": id})
}

func DeleteItem(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    id := c.Param("id")
    result, err := db.Exec("DELETE FROM ITEM WHERE id = ?", id)
    if err != nil {
        fmt.Errorf("DeleteItem: %v", err)
    }
    id_deleted, err := result.RowsAffected()
    if err != nil {
        fmt.Errorf("DeleteItem: %v", err)
    }
    c.IndentedJSON(http.StatusCreated, gin.H{ "id": id_deleted})
}

func GetUserItems(c *gin.Context) {
    var items []Item
    id := c.Param("id")
    rows, err := db.Query("SELECT * FROM ITEM WHERE owner_id = ?", id)
    if err != nil {
        fmt.Errorf("GetUserItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.Id, &item.Owner_id, &item.Name, &item.Size, &item.Link); err != nil {
            fmt.Errorf("GetUserItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        fmt.Errorf("GetUserItems: %v", err)
    }
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.IndentedJSON(http.StatusOK, items)
}

func GetAllUserItems(id int64) ([]Item, error){
    var items []Item
    rows, err := db.Query("SELECT * FROM ITEM WHERE owner_id = ?", id)
    if err != nil {
        return nil, fmt.Errorf("GetAllUserItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.Id, &item.Owner_id, &item.Name, &item.Size, &item.Link); err != nil {
            return nil, fmt.Errorf("GetAllUserItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetAllUserItems: %v", err)
    }
    return items, nil
}

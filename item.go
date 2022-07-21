package main

import (
    "fmt"
)

type Item struct {
    id          int64
    owner_id    int64
    name        string
    size        string
    link        string
}

func GetItems() ([]Item, error){
    var items[]Item
    rows, err := db.Query("SELECT * FROM ITEM")
    if err != nil {
        return nil, fmt.Errorf("GetItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.id, &item.owner_id, &item.name, &item.size, &item.link); err != nil {
            return nil, fmt.Errorf("GetItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetItems: %v", err)
    }
    return items, nil
}

func AddItem(item Item, owner User) (int64, error) {
    result, err := db.Exec("INSERT INTO ITEM (owner_id, name, size, link) VALUES ((SELECT id FROM USER WHERE id = ?), ?, ?, ?)", owner.id, item.name, item.size, item.link)
    if err != nil {
        return 0, fmt.Errorf("AddItem: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddItem: %v", err)
    }
    return id, nil
}

func GetUserItems(owner User) ([]Item, error) {
    var items[]Item
    rows, err := db.Query("SELECT * FROM ITEM WHERE owner_id = ?", owner.id)
    if err != nil {
        return nil, fmt.Errorf("GetUserItems: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.id, &item.owner_id, &item.name, &item.size, &item.link); err != nil {
            return nil, fmt.Errorf("GetUserItems: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetUserItems: %v", err)
    }
    return items, nil
}

func GetUserItemsByID(id int64) ([]Item, error) {
    var items[]Item
    rows, err := db.Query("SELECT * FROM ITEM WHERE owner_id = ?", id)
    if err != nil {
        return nil, fmt.Errorf("GetUserItemsByID: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.id, &item.owner_id, &item.name, &item.size, &item.link); err != nil {
            return nil, fmt.Errorf("GetUserItemsByID: %v", err)
        }
        items = append(items, item)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("GetUserItemsByID: %v", err)
    }
    return items, nil
}

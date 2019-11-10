package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Tag struct {
    ID int
    Name string
}

func main() {
    fmt.Println("Go MySQL Tutorial")

    db, err := sql.Open("mysql", "root:***@tcp(127.0.0.1:3306)/***")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    insert, err := db.Query("INSERT INTO test VALUES (2, 'TEST2')")

    if err != nil {
        panic(err.Error())
    }

    defer insert.Close()

    results, err := db.Query("SELECT name FROM test")

    if err != nil {
        panic(err.Error())
    }
    var l [] string

    for results.Next(){
        // var tag Tag
        var name string
        // err = results.Scan(&tag.ID, &tag.Name)
        err = results.Scan(&name)
        if err != nil {
            panic(err.Error())
        }
        l = append(l, name)
        log.Println(l)
        log.Println(len(l))
    }
} 
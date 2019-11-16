package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    // "errors"

)

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
    results ,err := db.Query("select name from test")
    if err != nil {
        panic(err.Error())
    }
    var l [] string
    for results.Next(){
        var name string
        results.Scan(&name)
        l = append(l, name)
    }
    fmt.Fprintln(w, l)
}

func addHandler (w http.ResponseWriter, r *http.Request){
    addName := r.URL.Path[len("/view/"):]
    q := "INSERT INTO test (name) VALUES ('"+ addName +"')"
    insert, err := db.Query(q)

    if err != nil {
        panic(err.Error())
    }

    insert.Close()
    fmt.Fprintf(w, "%s added", addName)
}

func initDB(){
	var err error
	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
    db, err = sql.Open("mysql", "root:elad@tcp(10.0.0.18:3306)/elad")
	if err != nil {
		panic(err)
	}
}

func main() {
    http.HandleFunc("/view/", handler)
    http.HandleFunc("/add/", addHandler)
    initDB()
    log.Fatal(http.ListenAndServe(":8080", nil))
}
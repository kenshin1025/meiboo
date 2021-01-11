package main

import (
	"database/sql"
	"fmt"
	"log"
	"meiboo/api/meibo"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	fmt.Println("Starting server at 'http://localhost:8080'")

	//DBに接続する
	db, err := ConnectSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", test)
	http.HandleFunc("/meibo", func(w http.ResponseWriter, r *http.Request) {
		meibo.Meibo(w, r, db)
	})
	http.HandleFunc("/meibo/update", func(w http.ResponseWriter, r *http.Request) {
		meibo.Update(w, r, db)
	})
	http.ListenAndServe(":8080", nil)
}

func ConnectSQL() (*sql.DB, error) {
	dsn := "root:root@tcp(db:3306)/meiboo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	Token   string `json:"token"`
	Image   string `json:"image"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type getMeiboResponse struct {
	Meibo []Member `json:"meibo"`
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func meibo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT token, image, name, comment FROM member WHERE workspace_id = 1")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var data getMeiboResponse
		for rows.Next() {
			var member Member
			if err := rows.Scan(&member.Token, &member.Image, &member.Name, &member.Comment); err != nil {
				log.Fatal(err)
			}
			data.Meibo = append(data.Meibo, member)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		res, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
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
		meibo(w, r, db)
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

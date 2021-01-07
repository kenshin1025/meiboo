package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type member struct {
	Image   string `json:"img"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type getMeiboResponse struct {
	Meibo []member `json:"meibo"`
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func meibo(w http.ResponseWriter, r *http.Request) {
	var data getMeiboResponse
	data.Meibo = []member{{"jack", "kenshin", "1人目"}, {"jeane", "健心", "こんにちは。よろしくお願いします"}, {"jodi", "あああ", "ああああああああああああああああああああああああああああああああああああ"}}
	res, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	fmt.Println("Starting server at 'http://localhost:8080'")

	http.HandleFunc("/", test)
	http.HandleFunc("/meibo", meibo)
	http.ListenAndServe(":8080", nil)
}

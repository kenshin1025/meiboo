package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
}

type getMeiboResponse struct {
	Meibo []user `json:"meibo"`
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func meibo(w http.ResponseWriter, r *http.Request) {
	var data getMeiboResponse
	data.Meibo = []user{{"kenshin"}, {"健心"}, {"あああ"}, {"afasfhafa"}}
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

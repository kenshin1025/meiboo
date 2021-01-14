package meibo

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type getTagsResponse struct {
	Tags []Tag `json:"tags"`
}

func Tags(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	//タグ一覧取得
	case "GET":
		rows, err := db.Query("SELECT id, name FROM tag")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var data getTagsResponse

		for rows.Next() {
			var tag Tag
			if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
				log.Fatal(err)
			}
			data.Tags = append(data.Tags, tag)
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

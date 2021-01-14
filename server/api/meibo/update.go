package meibo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type updateRequest struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	case "PUT":
		// リクエストbody(json)を受け取る
		header := r.Header
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var request updateRequest
		err := json.Unmarshal(buf.Bytes(), &request)
		if err != nil {
			log.Fatal(err)
		}
		if request.Type == "name" {
			// トランザクション開始
			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			// tokenを元にユーザーのnameを更新
			_, execErr := tx.Exec("UPDATE member SET name = ? WHERE token = ?", request.Value, header.Get("member-token"))
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}
			// エラーが起きなければコミット
			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
			}

			// レスポンス
			w.WriteHeader(http.StatusOK)
		} else if request.Type == "comment" {
			// トランザクション開始
			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			// tokenを元にユーザーのnameを更新
			_, execErr := tx.Exec("UPDATE member SET comment = ? WHERE token = ?", request.Value, header.Get("member-token"))
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}
			// エラーが起きなければコミット
			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
			}

			// レスポンス
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
		}
	}
}

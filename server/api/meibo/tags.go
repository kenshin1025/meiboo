package meibo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type getTagsResponse struct {
	Tags []Tag `json:"tags"`
}

type postTagsRequest struct {
	TagName string `json:"tagName"`
}

type postTagsResponse struct {
	TagID int64 `json:"tagID"`
}

func Tags(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
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
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	case "POST":
		// リクエストbody(json)を受け取る
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var request postTagsRequest
		err := json.Unmarshal(buf.Bytes(), &request)
		if err != nil {
			log.Fatal(err)
		}

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		var execErr error
		// DBに追加
		//レコードを取得する必要のない、クエリはExecメソッドを使う。
		result, execErr := tx.Exec("INSERT INTO tag(name) VALUES(?)", request.TagName)
		//エラーが起きたらロールバック
		if execErr != nil {
			_ = tx.Rollback()
			log.Fatal(execErr)
		}

		// エラーが起きなければコミット
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		tagID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		data := postTagsResponse{TagID: tagID}

		res, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

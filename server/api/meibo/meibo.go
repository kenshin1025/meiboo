package meibo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"meiboo/middleware/token"
	"net/http"
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

type postMeiboRequest struct {
	Name string `json:"name"`
}

func Meibo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	switch r.Method {
	//名簿一覧取得
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

	//名簿にメンバー追加
	case "POST":
		// リクエストbody(json)を受け取る
		body := r.Body
		defer body.Close()

		// byte配列に変換するためにcopy
		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		// byte配列にしたbody内のjsonをgoで扱えるようにobjectに変換
		var request postMeiboRequest
		err := json.Unmarshal(buf.Bytes(), &request)
		if err != nil {
			log.Fatal(err)
		}

		// tokenとしてuuid作成
		token, err := token.CreateToken()
		if err != nil {
			log.Fatal(err)
		}

		// トランザクション開始
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// DBに追加
		//レコードを取得する必要のない、クエリはExecメソッドを使う。
		_, execErr := tx.Exec("INSERT INTO member(token, name, workspace_id, image, comment) VALUES(?,?,?,?,?)", token, request.Name, 1, "", "")
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
		w.WriteHeader(http.StatusOK)
	}
}

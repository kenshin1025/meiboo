package meibo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"meiboo/middleware/token"
	"net/http"
	"strconv"
	"strings"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Member struct {
	Token   string `json:"token"`
	Image   string `json:"image"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Tags    []Tag  `json:"tags"`
}

type getMeiboResponse struct {
	Meibo []Member `json:"meibo"`
}

type postMeiboRequest struct {
	Image   string `json:"image"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Tags    []Tag  `json:"tags"`
}

type postMeiboResponse struct {
	Token string `json:"token"`
}

func Meibo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	//名簿一覧取得
	case "GET":
		rows, err := db.Query("SELECT token, image, name, comment, tagNames, tagIDs FROM (SELECT * FROM member WHERE workspace_id = 1) AS T1 JOIN (select member_tag.member_id, GROUP_CONCAT(tag.id) AS tagIDs, GROUP_CONCAT(tag.name) AS tagNames from member_tag join tag ON member_tag.tag_id = tag.id GROUP BY member_tag.member_id) AS T2 ON T1.id = T2.member_id")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var data getMeiboResponse
		for rows.Next() {
			var member Member
			var tagNames string
			var tagIDs string
			if err := rows.Scan(&member.Token, &member.Image, &member.Name, &member.Comment, &tagNames, &tagIDs); err != nil {
				log.Fatal(err)
			}
			sliceTagNames := strings.Split(tagNames, ",")
			sliceTagIDs := strings.Split(tagIDs, ",")
			for i := range sliceTagIDs {
				tagID, err := strconv.Atoi(sliceTagIDs[i])
				if err != nil {
					log.Fatal(err)
				}
				tagName := sliceTagNames[i]
				member.Tags = append(member.Tags, Tag{tagID, tagName})
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

		var execErr error
		// DBに追加
		//レコードを取得する必要のない、クエリはExecメソッドを使う。
		result, execErr := tx.Exec("INSERT INTO member(token, name, image, comment, workspace_id) VALUES(?,?,?,?,?)", token, request.Name, request.Image, request.Comment, 1)
		//エラーが起きたらロールバック
		if execErr != nil {
			_ = tx.Rollback()
			log.Fatal(execErr)
		}

		if len(request.Tags) >= 1 {
			memberID, err := result.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}

			var insertValuesString string
			for i, v := range request.Tags {
				if i != 0 {
					insertValuesString += ","
				}
				insertValuesString += "(" + strconv.FormatInt(memberID, 10) + "," + strconv.Itoa(v.ID) + ")"
			}
			insertValuesString += ";"

			query := "INSERT INTO member_tag(member_id, tag_id) VALUES" + insertValuesString

			_, execErr := tx.Exec(query)
			//エラーが起きたらロールバック
			if execErr != nil {
				_ = tx.Rollback()
				log.Fatal(execErr)
			}
		}

		// エラーが起きなければコミット
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		data := postMeiboResponse{token}

		res, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

# meiboo
名簿アプリ（情報系組織がメイン）を作成する\
いつでもこれを見れば、組織にいる人の今のことが知れることを目指す

## ミッション
組織に属する人のことを知れる

## 解決したい課題
新しく組織に所属した人が自分が属す組織にどんな人がいるのか、顔と名前が一致しない、自分が質問をしたいことは誰に聞けばいいのかわからない

## 使用技術
* Nuxt.js
  * フロント
* Golang
  * バック
* MySQL
  * データベース
* Docker
  * 開発環境
* Firebase
  * 認証

## 要件
* 属している人のことが一覧で見られる
  * 画像
  * 名前
  * GitHub
* サインアップ、サインイン可能
* 自分のプロフィールを編集できる

## 自分が作ろうとしているものの問題
* オリジナリティあるか？
  * 別にこれを使わなくても組織で名簿って作れるんじゃない？
    * 紙媒体
    * 自作サイト
    * エクセル
    * それらに比べて編集しやすい、作りやすい、見やすい?
* 最低限の実装になっている
  * プラスアルファ何ができるか
* デプロイどうやってするの？
  * フロントはFirebase
  * バックエンドはどうする？
  * とりあえずローカルで動けばいい


## コメント
* 自由度が高い
* タグ付けができる
* GitHubAPI利用
* メンバーが多い、チームが色々変わる組織ほど有効
* 
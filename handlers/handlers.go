package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/gorilla/mux"
)

// /helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// デコーダの導入
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// TODO:リクエストボディから取得したデータをDBに挿入して、実装にデータベースに格納した値を得る
	article := reqArticle

	// エンコーダの導入
	json.NewEncoder(w).Encode(article)
}

// /article/listのハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	// パラメータのpageが1つ以上あるなら
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		// 数値に変換できないのであれば400エラーを出す
		if err != nil {
			errMsg := fmt.Sprintf("Invalid query parameter (page %d)", page)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
	} else {
		// pageが存在しない場合
		page = 1
	}

	// TODO: 記事一覧をデータベースから取得する
	articleList := []models.Article{models.Article1, models.Article2}
	// エンコード
	json.NewEncoder(w).Encode(articleList)
}

// /article/1のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		errString := fmt.Sprintf("Invalid query parameter (articleID %d)", articleID)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}
	// TODO: サービス層の機能：指定IDの記事をDBから取得する
	article := models.Article1
	// エンコード
	json.NewEncoder(w).Encode(article)
}

// /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}
	// TODO:指定した記事にいいね+1する更新作業をDBに保存する
	json.NewEncoder(w).Encode(article)
}

// commentのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}
	// TODO: DBを挿入して、実際DBに格納された値を得る
	comment := reqComment

	json.NewEncoder(w).Encode(comment)
}

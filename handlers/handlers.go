package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/Takao-Yamasaki/myapi_v2/services"
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

	article, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

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

	articleList, err := services.GetArticleListService(page)
	log.Printf("handler: %v", articleList)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

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

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}
	// エンコード
	json.NewEncoder(w).Encode(article)
}

// /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}
	article, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// commentのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json", http.StatusBadRequest)
		return
	}

	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}

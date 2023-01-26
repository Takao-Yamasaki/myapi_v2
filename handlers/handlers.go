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
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusBadRequest)
		return
	}
	w.Write(jsonData)
}

// /article/1のハンドラ
func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		errString := fmt.Sprintf("Invalid query parameter (articleID %d)", articleID) 
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusBadRequest)
		return
	}
	w.Write(jsonData)
}

// /article/niceのハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusBadRequest)
	}
	w.Write(jsonData)
}

// commentのハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json", http.StatusBadRequest)
	}
	w.Write(jsonData)
}

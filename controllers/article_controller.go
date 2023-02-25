package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Takao-Yamasaki/myapi_v2/apperrors"
	"github.com/Takao-Yamasaki/myapi_v2/controllers/services"
	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/gorilla/mux"
)

// Article用サービスインターフェース
type ArticleController struct {
	service services.ArticleServicer
}

// コンストラクタ関数
func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// /helloのハンドラ
func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /articleのハンドラ
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	// デコーダの導入
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad reqest body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// エンコーダの導入
	json.NewEncoder(w).Encode(article)
}

// /article/listのハンドラ
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	// パラメータのpageが1つ以上あるなら
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		// 数値に変換できないのであれば400エラーを出す
		if err != nil {
			err = apperrors.BadPathParam.Wrap(err, "pathparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		// pageが存在しない場合
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	// エンコード
	json.NewEncoder(w).Encode(articleList)
}

// /article/1のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadPathParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	// エンコード
	json.NewEncoder(w).Encode(article)
}

// /article/niceのハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		apperrors.ErrorHandler(w, req, err)
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

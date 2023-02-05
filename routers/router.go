package routers

import (
	"net/http"

	"github.com/Takao-Yamasaki/myapi_v2/controllers"
	"github.com/gorilla/mux"
)

// パスとハンドラ関数の対応づけがされたgorilla/muxルータを作成して、戻り値を返却する
func NewRouter(con *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()

	// r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	// 対応づけが完了したルータを返す
	return r
}

package api

import (
	"database/sql"
	"net/http"

	"github.com/Takao-Yamasaki/myapi_v2/controllers"
	"github.com/Takao-Yamasaki/myapi_v2/services"
	"github.com/gorilla/mux"
)

// パスとハンドラ関数の対応づけがされたgorilla/muxルータを作成して、戻り値を返却する
func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	r := mux.NewRouter()

	// r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	// 対応づけが完了したルータを返す
	return r
}

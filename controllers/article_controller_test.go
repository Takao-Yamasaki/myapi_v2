package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticleListHandler(t *testing.T) {
	// 2.テスト対象のハンドラメソッドに入れるinputを定義
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: ハンドラに渡す２つの引数
			// w http.ResponseWriter, req *http.Request を用意する
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			// 3.テスト対象のハンドラメソッドからoutputを得る
			res := httptest.NewRecorder()
			// ユニットテストの中で、具体的に引数に何を渡せばいいのかがわからない
			aCon.ArticleListHandler(res, req)

			// 4.outputが期待通りかチェック
			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

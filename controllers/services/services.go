// インターフェースを定義する
// MyAppService構造体という具体的な型ではなく、次のメソッドを持つインターフェースである
package services

import "github.com/Takao-Yamasaki/myapi_v2/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Article, error)
}

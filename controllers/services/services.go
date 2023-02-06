// インターフェースを定義する
// MyAppService構造体という具体的な型ではなく、次のメソッドを持つインターフェースである
package services

import "github.com/Takao-Yamasaki/myapi_v2/models"

// article関連を引き受けるサービス
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// comment関連を引き受けるサービス
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}

package services

import (
	"go/doc/comment"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/Takao-Yamasaki/myapi_v2/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, nil
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList := append(article.CommentList, commentList...)
	
	return article, nil
}

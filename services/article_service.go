/*
サービス関数の実装
*/
package services

import (
	_ "fmt"
	"log"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/Takao-Yamasaki/myapi_v2/repositories"
)

// 記事データをDB内に挿入し、その値を返す
func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, nil
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, nil
	}
	return newArticle, nil
}

// クエリパラメータで指定したページの記事一覧をDBから取得する
func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	log.Printf("service: %v", articleList)
	if err != nil {
		return []models.Article{}, nil
	}
	return articleList, nil
}

// 指定IDの記事をDBから取得する
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

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// 指定IDの記事のいいね数を+1して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, nil
	}
	defer db.Close()

	if err = repositories.UpdateNiceNum(db, article.ID); err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Contents:    article.Contents,
		UserName:    article.UserName,
		NiceNum:     article.NiceNum + 1,
		CommentList: article.CommentList,
		CreatedAt:   article.CreatedAt,
	}, nil
}

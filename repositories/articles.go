/*
リポジトリ層
*/
package repositories

import (
	"database/sql"
	"log"

	"github.com/Takao-Yamasaki/myapi_v2/models"
)

const (
	articleNumPerPage = 5
)

// 記事を挿入する関数
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	log.Println("repositories: start")
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			log.Println(err)
			return []models.Article{}, err
		}
		articleArray = append(articleArray, article)
	}
	log.Printf("repositories: %v", articleArray)
	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var article models.Article
	if err = row.Scan(&article.NiceNum); err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, article.NiceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

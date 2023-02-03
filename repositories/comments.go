package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	newComment.CommentID = int(id)
	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)

	for rows.Next() {
		var createdTime sql.NullTime
		var comment models.Comment

		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}

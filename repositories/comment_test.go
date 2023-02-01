package repositories_test

import (
	"testing"

	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/Takao-Yamasaki/myapi_v2/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d, but got ID %d", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "insertTest",
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}

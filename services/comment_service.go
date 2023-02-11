package services

import (
	"github.com/Takao-Yamasaki/myapi_v2/apperrors"
	"github.com/Takao-Yamasaki/myapi_v2/models"
	"github.com/Takao-Yamasaki/myapi_v2/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}
	return newComment, nil
}

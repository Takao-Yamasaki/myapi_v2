package testdata

import "github.com/Takao-Yamasaki/myapi_v2/models"

// サービスモック構造体
type serviceMock struct{}

// サービスモック構造体のコンストラクタ
func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

// 5つのメソッド
func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}

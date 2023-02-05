package services

import "database/sql"

// サービス構造体の定義
type MyAppService struct {
	db *sql.DB
}

// コンストラクタの定義
func NewMyAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}

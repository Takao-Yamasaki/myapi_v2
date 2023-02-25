package apperrors

type MyAppError struct {
	// TODO: 独自エラーに含めるフィールドの定義
	// フィールド名を省略した場合、型名がそのままフィールド名になる
	ErrCode        // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err     error  // エラーチェーンのための内部エラー
}

// 根本のエラー原因を表すErrフィールドの内容を出力
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

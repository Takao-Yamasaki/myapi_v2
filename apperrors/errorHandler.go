package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生した時のレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// 受け取ったエラーを独自エラー型に変換する
	var appErr *MyAppError
	// errors.As関数で引数のerrをMyAppErr型のappErrに変換
	if !errors.As(err, &appErr) {
		// もし変換に失敗したら、Unknownエラーを変数appErrに手動で格納
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}

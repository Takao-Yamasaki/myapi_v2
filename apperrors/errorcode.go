package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000" // 開発者が想定しなかったエラー

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002" // Select文の実行に失敗したときのエラーコード
	NAData           ErrCode = "S003" // 指定された記事がなかった
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
)

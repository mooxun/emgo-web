package ecode

const (
	Success = 0
	Error   = 600
	// 验证
	AccountException       = 400
	Unauthorized           = 401
	LoginFailed            = 402
	PermissionDenied       = 403
	PageNotFound           = 404
	ValidatesRequestsError = 414
	// 查询
	ResultEmpty    = 700
	OrmUpdateError = 701
	OrmCreateError = 702
	ParamsError    = 707
)

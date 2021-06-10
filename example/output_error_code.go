package example

type ErrCode int

const (
	ERROR_REASON_UNSPECIFIED ErrCode = 0 // unknown
	USER_NOT_FOUND ErrCode = 1 // not found
	Role_Create_Fail ErrCode = 300001001 // 新增角色失败

)

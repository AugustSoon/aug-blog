package errno

var (
	OK                  = &Errno{Code: 0, Message: "ok"}
	InternalServerError = &Errno{Code: 10001, Message: "服务"}
	ErrBind             = &Errno{Code: 10002, Message: "参数错误"}

	ErrValidation   = &Errno{Code: 20001, Message: "参数验证失败"}
	ErrDatabase     = &Errno{Code: 20002, Message: "数据库错误"}
	ErrToken        = &Errno{Code: 20003, Message: "生成token失败"}
	ErrTokenInvalid = &Errno{Code: 20004, Message: "未登录或登录超时"}

	ErrEncrypt      = &Errno{Code: 20101, Message: "密码加密失败"}
	ErrUserNotFound = &Errno{Code: 20102, Message: "用户不存在"}
	ErrUserExist    = &Errno{Code: 20103, Message: "用户已经存在"}
	ErrLogin        = &Errno{Code: 20104, Message: "账号或密码不正确"}
)

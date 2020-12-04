package exception

type Exception struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewParamException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 4000
	exception.Msg = "Invalidated parameter"
	if len(msg) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}

func NewDataTypeException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 4001
	exception.Msg = "请求类型不存在"
	if len(msg) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}


func NewUncatchedException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 6000
	exception.Msg = "System error - uncatched exception"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}

func NewValidationException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 2
	exception.Msg = "Item size has exceeded the maximum allowed size"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}



func NewDBException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 4
	exception.Msg = "Mysql error"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}

func NewInnerException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5
	exception.Msg = "Inner error"
	if len(msg) > 0 && len(msg[0]) > 0 {
		exception.Msg = msg[0]
	}
	return exception
}



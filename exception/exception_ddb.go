package exception

const (
	DdbCCF = "ConditionalCheckFailedException"        //有条件请求失败
	DdbRNF = "ResourceNotFoundException"                          //ResourceNotFoundException 未找到请求的资源
	DdbT   = "ThrottlingException"                    //请求速率超出吞吐量上限
	DdbPTE = "ProvisionedThroughputExceededException" //已超过表或一个或更多全局二级索引的最大允许预置吞吐量
	DdbRLE = "RequestLimitExceeded"                   //吞吐量超出您的账户的当前吞吐量限制

	DdbNoItem = "no item found" //ResourceNotFoundException未找到请求的资源
)

func DdbNoItemFound(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5000
	exception.Msg = "不存在。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbRNFException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5000
	exception.Msg = "不存在。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbCCFException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5001
	exception.Msg = "保存异常，请重试。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbTException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5002
	exception.Msg = "请求速率超出吞吐量上限。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbPTEException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5003
	exception.Msg = "超出最大允许预置吞吐量。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbRLFException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5004
	exception.Msg = "吞吐量超出您的账户的当前吞吐量限制。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

func DdbOtherException(msg ...string) Exception {
	exception := Exception{}
	exception.Code = 5005
	exception.Msg = "dynamo db 异常。"
	if len(msg) > 0 {
		exception.Msg += " 异常信息: " + msg[0]
	}
	return exception
}

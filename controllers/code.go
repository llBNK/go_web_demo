package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:    "success",
	CodeServerBusy: "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}

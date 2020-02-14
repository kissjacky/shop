package controller

type Controller struct {
}

type retOk struct {
	Errno int         `json:"errno"`
	Data  interface{} `json:"data"`
}
type retErr struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
	Field string `json:"field"`
}

func RetOk(data interface{}) *retOk {
	return &retOk{0, data}
}
func RetErr(errno int, msg string, field string) *retErr {
	return &retErr{errno, msg, field}
}

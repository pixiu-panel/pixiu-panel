package response

// Result
// @description: 响应
// @receiver r
// @param code int 状态码
// @param data any 数据
// @param msg string 消息
// @param err string 错误信息
// @return err error 返回数据错误
func (r *Response) Result() {
	type resp struct {
		Code   int    `json:"code"`
		Data   any    `json:"data"`
		Msg    string `json:"message"`
		ErrMsg string `json:"errMsg,omitempty"`
	}

	rd := resp{
		r.code,
		r.data,
		r.msg,
		r.errMsg,
	}
	// 返回数据
	r.ctx.JSON(r.code, rd)
}

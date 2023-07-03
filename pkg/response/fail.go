package response

// Fail
// @description: 失败响应
// @receiver r
// @param data
// @return err
func (r *Response) Fail() {
	if r.msg == "" {
		r.msg = "系统错误"
	}
	if r.code == 0 {
		r.code = fail
	}
	r.Result()
}

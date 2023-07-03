package response

// Success
// @description: 成功响应
// @receiver r
// @param data
// @return err
func (r *Response) Success() (err error) {
	if r.msg == "" {
		r.msg = "success"
	}
	if r.code == 0 {
		r.code = success
	}
	return r.Result()
}

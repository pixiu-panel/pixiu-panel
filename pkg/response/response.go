package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// 定义状态码
const (
	fail    = http.StatusInternalServerError
	success = http.StatusOK
)

// Response
// @description: 返回结果
type Response struct {
	ctx    echo.Context
	code   int
	data   any
	msg    string
	errMsg string
}

// New
// @description: 返回结果实现
// @param ctx
// @return Response
func New(ctx echo.Context) *Response {
	var r Response
	r.ctx = ctx

	return &r
}

// SetCode
// @description: 设置状态码
// @receiver r
// @param code
// @return *Response
func (r *Response) SetCode(code int) *Response {
	r.code = code
	return r
}

// SetData
// @description: 设置返回数据
// @receiver r
// @param data
// @return *Response
func (r *Response) SetData(data any) *Response {
	r.data = data
	return r
}

// SetMsg
// @description: 设置返回消息
// @receiver r
// @param msg
// @return *Response
func (r *Response) SetMsg(msg string) *Response {
	r.msg = msg
	return r
}

// SetError
// @description: 设置错误信息
// @receiver r
// @param err
// @return *Response
func (r *Response) SetError(err error) *Response {
	if err != nil {
		r.errMsg = err.Error()
	}
	return r
}

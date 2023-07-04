package model

// Response
// @description: 通用响应结构体
type Response[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

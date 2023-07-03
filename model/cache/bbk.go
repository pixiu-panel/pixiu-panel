package cache

// BBKBaseResponse
// @description: BBK接口返回基础结构体
type BBKBaseResponse[T any] struct {
	Code     int    `json:"code"`     // 错误码
	Msg      string `json:"msg"`      // 错误信息
	Data     T      `json:"data"`     // 数据
	ErrorMsg string `json:"errorMsg"` // 错误信息
}

type BBKJdQrcode struct {
	Qr      string `json:"qr"`      // 二维码 Base64
	QrUrl   string `json:"qrUrl"`   // 扫码登录地址
	Timeout int    `json:"timeout"` // 有效期(秒)
	Cookie  string `json:"cookie"`  // Cookie数据，手动填充
}

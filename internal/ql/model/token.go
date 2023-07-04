package model

import "encoding/json"

// Token
// @description: 青龙 Token 数据
type Token struct {
	Token      string `json:"token"`
	TokenType  string `json:"token_type"`
	Expiration int    `json:"expiration"`
}

// MarshalBinary
// @description: Redis序列化
// @receiver s
// @return []byte
// @return error
func (s *Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// UnmarshalBinary
// @description: Redis反序列化
// @receiver s
// @param b
// @return error
func (s *Token) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, s)
}

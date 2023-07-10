package tests

import (
	"pixiu-panel/internal/bbk"
	"testing"
)

// TestGetJdQrcode
// @description: 获取京东二维码
// @param t
func TestGetJdQrcode(t *testing.T) {
	qr, err := bbk.GetJdQrcode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("二维码获取成功")
	t.Logf("二维码地址: %s", qr.QrUrl)
	t.Logf("二维码有效期: %d", qr.Timeout)
	t.Logf("二维码 Base64: %s", qr.Qr)
	t.Logf("Cookie: %s", qr.Cookie)
}

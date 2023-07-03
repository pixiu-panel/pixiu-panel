package jd

import (
	"github.com/go-resty/resty/v2"
	"pixiu-panel/model/cache"
)

// GetUserInfo
// @description: 获取用户信息
// @param cookie string cookie
// @return data cache.JdBaseResponse[cache.JdUserInfo] 用户信息
// @return err error 错误信息
func GetUserInfo(cookie string) (data cache.JdUserInfo, err error) {
	headers := map[string]string{
		"Accept":          "*/*",
		"Connection":      "keep-alive",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.62",
		"Host":            "me-api.jd.com",
		"Accept-Language": "zh-cn",
		"Referer":         "https://home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&",
		"Accept-Encoding": "deflate, br",
		"Cookie":          cookie,
	}

	var resp cache.JdBaseResponse[cache.JdUserInfo]

	cli := resty.New()
	_, err = cli.R().
		SetHeaders(headers).
		SetResult(&resp).
		Get("https://me-api.jd.com/user_new/info/GetJDUserInfoUnion")
	if err != nil {
		return
	}
	data = resp.Data
	return
}

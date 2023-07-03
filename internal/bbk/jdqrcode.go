package bbk

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
	"pixiu-panel/model/cache"
	"strings"
	"time"
)

// GetJdQrcode
// @description: 获取京东二维码
// @return data cache.BBKBaseResponse[cache.BBKJdQrcode] 二维码信息
// @return err error 错误信息
func GetJdQrcode() (data cache.BBKJdQrcode, err error) {
	if config.Conf.BBK.JdQr.Host == "" {
		err = fmt.Errorf("未配置京东二维码服务地址")
		return
	}
	// 组装请求地址
	api := fmt.Sprintf("%s/d/getQR?t=%d", config.Conf.BBK.JdQr.Host, time.Now().Local().UnixMilli())

	var respData cache.BBKBaseResponse[cache.BBKJdQrcode]

	cli := resty.New()
	resp, err := cli.R().
		SetResult(&respData).
		Get(api)
	if err != nil {
		return
	}

	// 从响应中取出数据
	data = respData.Data
	// 手动补充 cookie
	data.Cookie = strings.TrimSpace(strings.Split(resp.Header().Get("Set-Cookie"), ";")[0])
	return
}

// CheckJdQrcode
// @description: 检查京东二维码是否扫描
// @param cookie string 从BBK获取二维码时的cookie
func CheckJdQrcode(cookie string) (data cache.BBKBaseResponse[string], err error) {
	// 组装请求地址
	api := fmt.Sprintf("%s/d/status?t=%d", config.Conf.BBK.JdQr.Host, time.Now().Local().UnixMilli())

	cli := resty.New()
	_, err = cli.R().
		SetHeader("Cookie", cookie).
		SetResult(&data).
		Get(api)
	if err != nil {
		return
	}

	return
}

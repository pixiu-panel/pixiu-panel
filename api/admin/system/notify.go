package system

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/config"
	"pixiu-panel/pkg/response"
	"reflect"
	"strings"
)

// GetNotifyConfig
// @description: 获取推送渠道配置
// @param ctx
func GetNotifyConfig(ctx *gin.Context) {
	// 手动组装一下返回数据
	var data = make(map[string]map[string]any)
	//data["wechat"] = map[string]any{
	//	"enable": config.Conf.Notify.Wechat.Enable,
	//	"qrcode": config.Conf.Notify.Wechat.QrCode,
	//}
	//data["qq"] = map[string]any{
	//	"enable": config.Conf.Notify.QQ.Enable,
	//	"qrcode": config.Conf.Notify.QQ.QrCode,
	//}
	v := reflect.ValueOf(config.Conf.Notify)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		if f.Kind() == reflect.Struct {
			// 是结构体，进行取出是否启用
			field := v.Type().Field(i).Name
			field = strings.ToLower(field)
			log.Debugf("当前处理的字段：%s", field)

			fieldType := v.Type().Field(i)
			if fieldType.Type.Kind() == reflect.Struct {
				// 是否启用
				enableField := f.FieldByName("Enable")
				if enableField.IsValid() && enableField.Kind() == reflect.Bool {
					data[field] = map[string]any{
						"enable": enableField.Bool(),
					}
				}
				// 二维码(未必都有)
				qrCodeField := f.FieldByName("QrCode")
				if qrCodeField.IsValid() && qrCodeField.Kind() == reflect.String {
					data[field]["qrcode"] = qrCodeField.String()
				}
			}
		}
	}

	response.New(ctx).SetData(data).Success()
}

package jd

import (
	"encoding/json"
	"net/url"
	"pixiu-panel/internal/bbk"
	cacheCli "pixiu-panel/internal/cache"
	"pixiu-panel/model/cache"
	"pixiu-panel/model/param"
	"pixiu-panel/service/jd"
	"regexp"
)

// checkJdBindStatus
// @description: 异步检查京东绑定状态
// @param key
func checkJdBindStatus(key string) {
	// 取出缓存的数据
	var che map[string]any
	cbs, err := cacheCli.Cache.Get([]byte(key))
	if err != nil {
		return
	}
	// 解析为map
	_ = json.Unmarshal(cbs, &che)

	for {
		// 查询二维码扫描状态
		var status cache.BBKBaseResponse[cache.BBKJdQrcodeScan]
		status, err = bbk.CheckJdQrcode(che["cookie"].(string))
		if err != nil {
			che["status"] = -1 // 标记为已过期
			break
		}
		// 是否需要更新缓存
		needUpdateCache := false
		// 是否需要退出循环
		flag := false
		// 处理检查结果
		switch status.Code {
		case 500, 202, 408:
			// 二维码失效
			che["status"] = -1 // 标记为已过期
			needUpdateCache = true
			flag = true
		case 200:
			// 还没扫描
		case 201:
			// 请在手机上确认登录
			che["status"] = 1 // 标记为已扫描待确认
			needUpdateCache = true
		case 410:
			// 登录成功
			// 提取出用户的PIN，准备入库
			pinMatch := regexp.MustCompile(`\[(.*?)\]`).FindStringSubmatch(status.Data.Msg)
			if len(pinMatch) != 2 {
				che["status"] = 3 // 标记为绑定失败
			} else {
				// 缓存用户PIN(URL 编码处理一下，防止中文乱码)
				che["pin"], _ = url.QueryUnescape(pinMatch[1])
				err = saveToDb(che["userId"].(string), che["pin"].(string))
				if err != nil {
					che["status"] = 3 // 标记为绑定失败
				} else {
					che["status"] = 2 // 标记为已绑定
				}
			}
			needUpdateCache = true
			flag = true
		}
		if needUpdateCache {
			// 更新缓存
			cacheDataBytes, _ := json.Marshal(che)
			_ = cacheCli.Cache.Set([]byte(key), cacheDataBytes, int(che["timeout"].(float64)))
		}
		if flag {
			break
		}
	}
}

// saveToDb
// @description: 保存到数据库
// @param userId
// @param pin
// @return err
func saveToDb(userId, pin string) (err error) {
	var p param.SaveJdAccount
	p.UserId = userId
	p.Pin = pin
	err = jd.SaveJdInfo(p)
	return
}

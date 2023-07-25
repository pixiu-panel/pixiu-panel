package tasks

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/jd"
	"pixiu-panel/model/cache"
	"pixiu-panel/model/entity"
	"strconv"
)

// updateJdAccountInfo
// @description: 更新京东账号信息
func updateJdAccountInfo() {
	// 查出所有未过期的账号
	var accounts []entity.UserJd
	err := db.Client.Where("expired = 0").Find(&accounts).Error
	if err != nil {
		log.Errorf("查询京东账号失败: %v", err)
		return
	}
	log.Debugf("共获取到 %d 个京东账号", len(accounts))
	// 调用京东接口查询基础信息并入库
	for _, account := range accounts {
		// 调用接口查询基础信息
		var info cache.JdUserInfo
		info, err = jd.GetUserInfo(account.Cookie)
		if err != nil {
			log.Errorf("获取京东账号 %s 基础信息失败: %v", account.Pin, err)
			continue
		}
		// 入库
		// 转换一下数据类型
		beanNum, _ := strconv.Atoi(info.AssetInfo.BeanNum)
		redBalance, _ := strconv.ParseFloat(info.AssetInfo.RedBalance, 64)
		couponNum, _ := strconv.Atoi(info.AssetInfo.CouponNum)
		pm := map[string]any{
			"nickname":    info.UserInfo.BaseInfo.Nickname,     // 更新昵称
			"avatar":      info.UserInfo.BaseInfo.HeadImageUrl, // 更新头像
			"level":       info.UserInfo.BaseInfo.LevelName,    // 更新等级
			"is_plus":     info.UserInfo.IsPlusVip,             // 更新是否是plus会员
			"bean_num":    beanNum,                             // 更新京豆数量
			"red_balance": redBalance,                          // 更新红包余额
			"coupon_num":  couponNum,                           // 更新优惠券数量
		}
		if err = db.Client.Model(&entity.UserJd{}).Where("pin = ?", account.Pin).Updates(pm).Error; err != nil {
			log.Errorf("更新京东账号 %s 基础信息失败: %v", account.Pin, err)
		}
	}
}

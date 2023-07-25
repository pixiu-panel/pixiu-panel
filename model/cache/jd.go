package cache

// JdBaseResponse
// @description: 京东接口返回基础结构体
type JdBaseResponse[T any] struct {
	Msg       string `json:"msg"`       // 错误信息
	Retcode   string `json:"retcode"`   // 错误码
	Timestamp int64  `json:"timestamp"` // 时间戳
	Data      T      `json:"data"`      // 数据
}

// JdUserInfo
// @description: 京东用户信息
type JdUserInfo struct {
	UserInfo struct {
		BaseInfo struct {
			Alias        string `json:"alias"`        // 京东账号
			CurPin       string `json:"curPin"`       // 原始 Id
			HeadImageUrl string `json:"headImageUrl"` // 头像地址
			LevelName    string `json:"levelName"`    // 等级名称
			Nickname     string `json:"nickname"`     // 昵称
			UserLevel    string `json:"userLevel"`    // 用户等级，不知道干嘛的
		} `json:"baseInfo"` // 基础信息
		IsPlusVip string `json:"isPlusVip"` // 是否是 Plus 会员
	} `json:"userInfo"` // 用户信息
	AssetInfo struct {
		BeanNum    string `json:"beanNum"`    // 京豆数量
		RedBalance string `json:"redBalance"` // 红包余额
		CouponNum  string `json:"couponNum"`  // 优惠券数量
	} `json:"assetInfo"` // 资产信息
}

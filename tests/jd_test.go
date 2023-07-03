package tests

import (
	"pixiu-panel/internal/jd"
	"testing"
)

// TestGetJdUserInfo
// @description: 获取京东用户信息
// @param t
func TestGetJdUserInfo(t *testing.T) {
	cookie := "pt_key=app_openAAJkohAQADDyCiTksyXuqTclxuV5SMIMqKAWK0Zj912i3sB1jejwCad8njk7VhJMiPDfUwAkJjU;pt_pin=zzsfmr;"
	ui, err := jd.GetUserInfo(cookie)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("用户信息获取成功")
	t.Logf("Id: %s", ui.Data.UserInfo.BaseInfo.CurPin)
	t.Logf("账号: %s", ui.Data.UserInfo.BaseInfo.Alias)
	t.Logf("昵称: %s", ui.Data.UserInfo.BaseInfo.Nickname)
	t.Logf("等级: %s", ui.Data.UserInfo.BaseInfo.LevelName)
	t.Logf("是否 Plus 会员: %v", ui.Data.UserInfo.IsPlusVip == "1")
}

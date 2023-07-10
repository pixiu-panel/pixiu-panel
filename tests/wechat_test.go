package tests

import (
	"encoding/json"
	"encoding/xml"
	"pixiu-panel/model/param"
	"testing"
)

// TestParseWechatAddFriend
// @description: 测试解析微信添加好友消息xml
// @param t
func TestParseWechatAddFriend(t *testing.T) {
	xmlStr := `{
    "content": "<msg fromusername=\"wxid_n5sg7tjup7f322\" encryptusername=\"v3_020b3826fd0301000000000032ca06c7fdcccb000000501ea9a3dba12f95f6b60a0536a1adb6a5fe9ca8a25b050f30e0e30bc9fcc9e7ff96448530544c2d5de9a50049aa72c9f430f891444ec8763b363ec84e609936fd11eab088f2934436abe890@stranger\" fromnickname=\"一个机器人\" content=\"HYXC\" fullpy=\"yigejiqiren\" shortpy=\"YGJQR\" imagestatus=\"3\" scene=\"30\" country=\"\" province=\"\" city=\"\" sign=\"\" percard=\"1\" sex=\"0\" alias=\"huayixinchen2020\" weibo=\"\" albumflag=\"0\" albumstyle=\"0\" albumbgimgid=\"\" snsflag=\"256\" snsbgimgid=\"\" snsbgobjectid=\"0\" mhash=\"15fd97506e235a415da6881a81ef9433\" mfullhash=\"15fd97506e235a415da6881a81ef9433\" bigheadimgurl=\"http://wx.qlogo.cn/mmhead/ver_1/6NoCXxylBD6orZmICqVRmnlpFQrODRxYUibtqWOFuCSvYHLvic67rEGM2LXFctHWvNFWWsPa5oqQmtzuVWlVtEiaLzycKADLP5y7LGLIy8HQHU/0\" smallheadimgurl=\"http://wx.qlogo.cn/mmhead/ver_1/6NoCXxylBD6orZmICqVRmnlpFQrODRxYUibtqWOFuCSvYHLvic67rEGM2LXFctHWvNFWWsPa5oqQmtzuVWlVtEiaLzycKADLP5y7LGLIy8HQHU/132\" ticket=\"v4_000b708f0b0400000100000000000f17e6f9521bb78c492efb9dab641000000050ded0b020927e3c97896a09d47e6e9e385d8ed81b8d51357c3bfcd804c08ae918d4c9fe7c65233bae08946e5dc5c90014c92f68a5bfca15e759a6c8599ebf9a6f8c2aedcacf389fe3ac06de2d4cf32de78368d7cf918bea5506d6d8b4fb5dea7d511f1490312900a40b21036b75930ec352b49aa511e595@stranger\" opcode=\"2\" googlecontact=\"\" qrticket=\"\" chatroomusername=\"\" sourceusername=\"\" sourcenickname=\"\" sharecardusername=\"\" sharecardnickname=\"\" cardversion=\"\" extflag=\"0\"><brandlist count=\"0\" ver=\"812848151\"></brandlist></msg>",
    "fromGroup": "fmessage",
    "fromUser": "fmessage",
    "time": "2023-07-10 13:58:19",
    "timestamp": 1688968699,
    "type": 37
}`
	// 解析基础结构
	var bd param.WechatCallback
	if err := json.Unmarshal([]byte(xmlStr), &bd); err != nil {
		t.Error(err)
		return
	}
	// 解析微信添加好友消息
	var addFriend param.WechatAddFriend
	if err := xml.Unmarshal([]byte(bd.Content), &addFriend); err != nil {
		t.Errorf("解析微信添加好友消息失败: %v", err)
		return
	}
	t.Logf("解析微信添加好友消息成功: %+v", addFriend)
}

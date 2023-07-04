package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"net/http"
	"pixiu-panel/internal/redis"
	"pixiu-panel/model/entity"
	userService "pixiu-panel/service/user"
	"pixiu-panel/utils"
	"time"
)

// LoginWithPassword
// @description: 账号密码登录模式
// @param _
// @param clientId
// @param userId
// @param password
// @return userID
// @return err
func LoginWithPassword(_ context.Context, clientId, username, password string) (userId string, err error) {
	log.Debugf("[%v]处理登录请求，账号：%s --> %s", clientId, username, password)

	// 取出用户信息
	var userInfo entity.User
	userInfo, err = userService.GetUserWithLogin(username)
	userId = userInfo.Id
	if err != nil {
		log.Errorf("获取用户信息失败: %v", err.Error())
		err = errors.New("账号不存在")
		return
	}
	// 校验密码
	if !utils.PasswordUtils().ComparePassword(userInfo.Password, password) {
		err = errors.New("密码错误")
		return
	}
	return
}

// ExtensionFields
// @description: 自定义响应Token的扩展字段
// @param _
// @return fieldsValue
func ExtensionFields(ti oauth2.TokenInfo) (fieldsValue map[string]any) {
	fieldsValue = map[string]any{}
	fieldsValue["license"] = "Made By Lixunhuan"
	fieldsValue["developer"] = "https://lxh.io"
	// 翻译一下token信息，方便前端使用(偷个懒，哈哈哈)
	fieldsValue["accessToken"] = ti.GetAccess()
	fieldsValue["refreshToken"] = ti.GetRefresh()
	fieldsValue["expires"] = time.Now().Local().Add(ti.GetAccessExpiresIn()).Format("2006/01/02 15:04:05")

	// 缓存用户的accessToken，为后续删除做准备
	key := fmt.Sprintf("oauth:token:token:%s", ti.GetUserID())
	if err := redis.Client.Set(context.Background(), key, ti.GetAccess(), time.Hour*24*30).Err(); err != nil {
		log.Errorf("缓存用户的accessToken失败: %v", err.Error())
	}

	// 还可以填充用户登录后需要的信息进去
	fieldsValue["username"] = ti.GetUserID()
	return
}

// ResponseToken
// @description: 返回Token生成结果
// @param w http.ResponseWriter 写入响应
// @param data map[string]any 响应数据
// @param header http.Header 响应头
// @param statusCode ...int 响应状态
// @return error 错误信息
func ResponseToken(w http.ResponseWriter, data map[string]any, header http.Header, statusCode ...int) error {
	log.Debugf("返回Token原始数据: %+v", data)
	type response struct {
		Code int            `json:"code"`
		Data map[string]any `json:"data"`
		Msg  string         `json:"message"`
	}

	status := http.StatusOK
	msg := "login success"
	if len(statusCode) > 0 && statusCode[0] > 0 {
		status = statusCode[0]
		msg = fmt.Sprintf("%v", data["error_description"])
		// 处理特殊返回 - 刷新Token到期了
		switch data["error"] {
		case "invalid_grant":
			msg = "登录已过期，请重新授权登录"
		case "invalid_request":
			msg = "登录参数错误"
		case "invalid_client":
			msg = "客户端验证失败"
		default:
			log.Errorf("收到未定义的登录错误: %v", data["error_description"])
		}
		data = nil
	}

	res := response{
		Code: status,
		Msg:  msg,
		Data: data,
	}

	jsonBytes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	for key := range header {
		w.Header().Set(key, header.Get(key))
	}

	w.WriteHeader(status)
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Errorf("返回Token失败: %v", err.Error())
		return err
	}
	return err
}

// InternalErrorHandler 自定义内部错误处理
func InternalErrorHandler(err error) (re *errors.Response) {
	re = errors.NewResponse(err, http.StatusUnauthorized)
	re.Description = err.Error()
	return
}

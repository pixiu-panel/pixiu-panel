package auth

import (
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	od "github.com/go-oauth2/redis/v4"
	"pixiu-panel/internal/redis"
	"pixiu-panel/pkg/auth/handle"
	"time"
)

var (
	OAuthServer *server.Server // 后台OAuth2服务
)

// InitOAuth2Server
// @description: 初始化后台OAuth2服务
func InitOAuth2Server() {
	manager := manage.NewDefaultManager()
	// 配置信息
	cfg := &manage.Config{
		AccessTokenExp:    time.Hour * 24,      // 访问令牌过期时间
		RefreshTokenExp:   time.Hour * 24 * 30, // 更新令牌过期时间
		IsGenerateRefresh: true,                // 是否生成新的更新令牌
	}
	// 设置密码模式的配置参数
	manager.SetPasswordTokenCfg(cfg)

	manager.MapTokenStorage(od.NewRedisStoreWithCli(redis.Client, "oauth:token:"))
	// 生成Token方式
	manager.MapAccessGenerate(handle.NewAccessGenerate())

	// 配置客户端
	// 配置客户端
	clientStore := store.NewClientStore()
	_ = clientStore.Set("admin", &models.Client{
		ID:     "admin",
		Secret: "kycgfS4sBjx2rPVD",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	// 设置密码登录模式处理逻辑
	srv.SetPasswordAuthorizationHandler(handle.LoginWithPassword)
	// 允许密码模式、刷新Token
	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	// 客户端ID和授权模式检查
	//srv.SetClientAuthorizedHandler(handle.CheckClient)
	// 自定义响应Token的扩展字段
	srv.SetExtensionFieldsHandler(handle.ExtensionFields)
	// 自定义返回数据接口
	srv.SetResponseTokenHandler(handle.ResponseToken)
	// 自定义内部错误处理
	srv.SetInternalErrorHandler(handle.InternalErrorHandler)

	OAuthServer = srv
}

package middleware

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixiu-panel/pkg/auth"
	"pixiu-panel/pkg/response"
	"strings"
)

// AuthorizeToken
// @description: 验证OAuth2生成的Token
// @return gin.HandlerFunc
func AuthorizeToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 判断有无token
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			response.New(ctx).SetMsg("请先登录").Fail()
			ctx.Abort()
			return
		}
		// 先取出用户Token
		token, err := auth.OAuthServer.ValidationBearerToken(ctx.Request)
		if err != nil {
			log.Errorf("获取Token失败，错误：%s", err.Error())
			response.New(ctx).SetCode(http.StatusUnauthorized).SetMsg("登录已失效或已在其他地方登录").Fail()
			ctx.Abort()
			return
		}

		// 判断通过，允许放行
		ctx.Set("userId", token.GetUserID())
		ctx.Next()
	}
}

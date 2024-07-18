package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/miao-crispy-corner/teamini/internal/pkg/core"
	"github.com/miao-crispy-corner/teamini/internal/pkg/errno"
	"github.com/miao-crispy-corner/teamini/internal/pkg/known"
	"github.com/miao-crispy-corner/teamini/internal/pkg/log"
)

// Auther 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Auther) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString(known.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method

		log.Debugw("Build authorize context", "sub", sub, "obj", obj, "act", act)
		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			core.WriteResponse(c, errno.ErrUnauthorized, nil)
			c.Abort()
			return
		}
	}
}

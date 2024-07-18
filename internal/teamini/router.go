package teamini

import (
	"github.com/gin-gonic/gin"
	"github.com/miao-crispy-corner/teamini/internal/teamini/controller/v1/user"
	"github.com/miao-crispy-corner/teamini/internal/teamini/store"

	"github.com/miao-crispy-corner/teamini/internal/pkg/core"
	"github.com/miao-crispy-corner/teamini/internal/pkg/errno"
	"github.com/miao-crispy-corner/teamini/internal/pkg/log"
	mw "github.com/miao-crispy-corner/teamini/internal/pkg/middleware"
	"github.com/miao-crispy-corner/teamini/pkg/auth"
)

// installRouters 安装 teamini 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 创建用户
			userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get) // 获取用户详情
		}
	}

	return nil
}

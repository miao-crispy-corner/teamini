package user

import (
	"github.com/gin-gonic/gin"

	"github.com/miao-crispy-corner/teamini/internal/pkg/core"
	"github.com/miao-crispy-corner/teamini/internal/pkg/errno"
	"github.com/miao-crispy-corner/teamini/internal/pkg/log"
	v1 "github.com/miao-crispy-corner/teamini/pkg/api/teamini/v1"
)

// Login 登录 teamini 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}

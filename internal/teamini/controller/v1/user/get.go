package user

import (
	"github.com/gin-gonic/gin"

	"github.com/miao-crispy-corner/teamini/internal/pkg/core"
	"github.com/miao-crispy-corner/teamini/internal/pkg/log"
)

// Get 获取一个用户的详细信息.
func (ctrl *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user function called")

	user, err := ctrl.b.Users().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

package auth

import (
	"time"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const (
	// casbin 访问控制模型.
	aclModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

// Authz 定义了一个授权器，提供授权功能.
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz 创建一个使用 casbin 完成授权的授权器.
func NewAuthz(db *gorm.DB) (*Authz, error) {
	// 初始化了一个 Gorm 适配器，adapter 将用于与数据库交互
	adapter, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	//模型定义了访问控制的结构和匹配规则
	m, _ := model.NewModelFromString(aclModel)

	// 创建了一个 SyncedEnforcer 对象 enforcer，它使用前面定义的模型 m 和 Gorm 适配器 adapter
	enforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}

	// 从数据库中加载策略。LoadPolicy() 方法会从数据库中读取策略
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	// 每 5s 从数据库中同步一次授权策略
	enforcer.StartAutoLoadPolicy(5 * time.Second)

	a := &Authz{enforcer}

	return a, nil
}

// Authorize 用来进行授权.
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}

package bootstrap

import (
	"github.com/go-ex/micro-omo/apps/admin/controllers"
	"github.com/go-ex/micro-omo/runtime/admin.proto"
)

// 加载启用的控制器
func loadController() {
	_ = admin.RegisterUserHandler(App.Server(), new(controllers.User))
}

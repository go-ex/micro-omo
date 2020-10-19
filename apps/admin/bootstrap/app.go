package bootstrap

import (
	"fmt"
	"github.com/go-ex/omo-cloud/db"
	micro "github.com/micro/go-micro/v2"
)

// 应用单例
var App = micro.NewService(
	micro.Name("go.micro.api.admin"),
)

func init() {
	// 加载配置等，在run前执行

	db.NewDatabase()
}

// 对外启动入口
func NewApp() {
	App.Init()

	loadController()

	// Run the server
	if err := App.Run(); err != nil {
		fmt.Println(err)
	}
}

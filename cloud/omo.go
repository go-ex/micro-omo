package cloud

// 项目全局服务容器
type AppBean interface {
	Singleton(alias string, fun interface{})
	Make(alias string) interface{}
}

// 服务提供者
type Providers interface {
	Construct(bean AppBean)
	// 注册启动调用
	Register() bool
	// 全部注册完后，调用boot
	Boot()
}

type Service interface {
}

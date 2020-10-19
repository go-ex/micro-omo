package main

import (
	admin "github.com/go-ex/micro-omo/apps/admin/bootstrap"
)

func main() {
	// 从应用的bootstrap开启启动
	admin.NewApp()
}

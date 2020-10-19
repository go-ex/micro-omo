# 安装
init:
	go get github.com/go-sql-driver/mysql
	go get github.com/micro/go-micro/v2
	go get github.com/micro/go-plugins/registry/consul/v2
	go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master


# 生成proto
proto:
	protoc --proto_path=./apps/admin/proto --go_out=./runtime/admin.proto --micro_out=./runtime/admin.proto ./apps/admin/proto/*.proto





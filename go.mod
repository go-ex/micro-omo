module github.com/go-ex/micro-omo

go 1.13

replace github.com/go-ex/omo-cloud => ./cloud

// https://github.com/etcd-io/etcd/issues/11563
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-ex/omo-cloud v1.18.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/http/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/transport/quic/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.2-0.20200728090142-c7f7e4a71077 // indirect
)

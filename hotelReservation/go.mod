module github.com/harlow/go-micro-services

go 1.13

require (
	github.com/akshayknarayan/burrito/resolv-go v0.0.0-00010101000000-000000000000
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/armon/go-metrics v0.3.0 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20171214222146-0e7658f8ee99
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/hashicorp/consul v1.0.6
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-rootcerts v0.0.0-20160503143440-6bb64b370b90 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/hashicorp/serf v0.8.1 // indirect
	github.com/mitchellh/go-homedir v0.0.0-20161203194507-b8bc1bf76747 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/olivere/grpc v1.0.0
	github.com/opentracing-contrib/go-stdlib v0.0.0-20180308002341-f6b9967a3c69
	github.com/opentracing/opentracing-go v1.0.2
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.11.2+incompatible
	github.com/uber/jaeger-lib v1.4.0 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.26.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

// TODO remove after publishing
replace github.com/akshayknarayan/burrito/resolv-go => ../../burrito/resolv-go

# go-shop-srvs

基于 go/grpc 的简单电商系统微服务, 该项目 web api 服务部分见  [go-shop-api](https://github.com/BetaCatPro/go-shop-api).

## Environment

- Golang 1.20.4
- protoc 23.2
- Consul 1.15.3
- Nacos 1.4.6
- ElasticSearch 7.x
- RocketMQ 4.5.1
- Jaeger 1.46.0

## Usage

1. Protobuf 生成 Go 代码

```shell
// 例如 user svr服务
protoc -I . user_srv/proto/user.proto --go_out=:. --go-grpc_out=require_unimplemented_servers=false:.
```

2. 启动 Consul 服务, 用于服务注册/注销/发现

```shell
consul agent -dev
```

3. 启动 Nacos 服务, 用于配置中心管理

```shell
// e.g. windows 环境
startup.cmd -m standalone
```

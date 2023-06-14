# go-shop-srvs

基于 go/grpc 的简单电商系统微服务, 该项目 web api 服务部分见  [go-shop-api](https://github.com/BetaCatPro/go-shop-api).

## Environment

基础环境

- Golang 1.20.4
- protoc 23.2
- Consul 1.15.3
- Nacos 1.4.6
- ElasticSearch 7.x
- Redis ==> 主要用于数据缓存、分布式锁
- PostgreSQL

分布式服务及网关工具

- 基于 RocketMQ 4.5.1 的提供可靠消息的最终一致性分布式事务方案
- 使用 Jaeger 1.46.0 作为链路追踪工具
- 使用 go-sentinel 作为限流，熔断工具，配置流量控制，熔断降级，系统保护等规则
- 使用 Kong 作为微服务 API 网关，配置身份认证、安全控制、流量控制等插件

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

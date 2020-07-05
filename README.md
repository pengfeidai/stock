# go-stock

## 快速开始

#### 代码仓库

```go
git clone git@github.com:pengfeidai/stock.git
```

#### 环境配置

- Go version >= 1.13
- Global environment configure

```go
export GO111MODULE=on
// 修改代理
export GOPROXY=https://goproxy.io
// go env -w GOPROXY=https://goproxy.cn,direct 
```

#### 服务启动

```go
cd go-stock

go run main.go

输出如下 `Listening and serving HTTP on Port: :3000, Pid: 15932`，表示 Http Server 启动成功。
```

#### 健康检查

```
curl -X GET http://127.0.0.1:3000/health_check?name=world
```

## 配置文件

```yaml
server:
  port: 3000
  mode: 'release'

mysql:
  user: root
  password: '123456'
  path: '127.0.0.1:3306'
  database: 'test'
  config: 'charset=utf8&parseTime=True&loc=Local'
  driver: 'mysql'
  maxIdleConns: 10
  maxOpenConns: 100
  log: true

log:
  debug: true
  maxAge: 7
  fileName: 'server.log'
  dirName: '/data/go-stock/logs'

url:
  # 路由前缀
  prefix: '/api/v1/go-stock'
```

### 启动容器

```shell
# docker build
make docker

# 启动容器（线上使用）
docker run --name go-stock --network=host \
  -v /opt/conf/go-stock:/opt/conf \
  -v /data/go-stock:/data/go-stock \
  -d go-stock
```
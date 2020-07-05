# 拉取golang镜像
# FROM golang:1.14.2-stretch
FROM golang:1.14.2-alpine3.11

# 设置镜像作者
LABEL MAINTAINER="302804389@qq.com"

ENV GOPROXY="https://goproxy.cn,direct" \
    GO111MODULE=on

# 创建目录,保存代码
RUN mkdir -p /opt/workspace/go-stock \
  && mkdir -p /data/go-stock/logs \
  && mkdir -p /opt/conf/

# 指定RUN、CMD与ENTRYPOINT命令的工作目录
WORKDIR /opt/workspace/go-stock

COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制所有文件到工作目录
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o go-stock .

COPY config.yaml /opt/conf

EXPOSE 3000

CMD ["./go-stock", "-c=/opt/conf/config.yaml"]
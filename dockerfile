FROM golang:1.15-alpine AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 拷贝源代码到镜像中
COPY . /go/src/Cloud-Native-Camp
# 编译
WORKDIR /go/src/Cloud-Native-Camp
RUN go install ./...

FROM alpine:3.13
COPY --from=builder /go/bin/Cloud-Native-Camp /bin/Cloud-Native-Camp
ENV ADDR=:8080

# 申明暴露的端口
EXPOSE 8080

# 设置服务入口
ENTRYPOINT [ "/bin/Cloud-Native-Camp" ]
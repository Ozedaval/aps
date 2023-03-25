# syntax=docker/dockerfile:1

# 从 docker hub 开始构建，速度慢。现采用从 ali hub 构建。如何切换到私有 hub
# 请参看：https://help.aliyun.com/document_detail/202437.html?spm=a2cl9.flow_devops2020_goldlog.0.0.32787b3asa6t33
#FROM golang:1.19-alpine
FROM golang:1.19-alpine
ENV GO111MODULE=on     GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /automation
EXPOSE 7201

CMD [ "/automation" ]

# 手动方式打包，并推送到 docker hub，需要配合第5行
# docker build --tag zhanghan2566/automation:0.1.2 .
# docker rPush zhanghan2566/automation:0.1.4

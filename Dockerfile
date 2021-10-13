FROM golang:1.13-alpine  as build

RUN echo -e https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.12/main/ > /etc/apk/repositories && apk add bash git

# Set the Current Working Directory inside the container
WORKDIR /app/wf-server

# We want to populate the module cache based on the go.{mod,sum} files.
ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./bin/mtd daemon/mtd/main.go


RUN git clone https://xiezhiqiang%40scrmtech.com:Waj33040!@e.coding.net/zqbc-scrm-new/scrm/go_conf.git ./go_conf


FROM alpine:3.12 as prod

# 设置时区为上海gca
RUN echo -e https://mirrors.tuna.tsinghua.edu.cn/alpine/v3.12/main/ > /etc/apk/repositories && apk update &&  apk add  tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" > /etc/timezone \
&& apk del tzdata


WORKDIR /app/wf-server

RUN mkdir -p ./runtime && chmod -R 777 ./runtime

COPY --from=build /app/wf-server/bin/ /app/wf-server/bin/
COPY  conf/ /app/wf-server/conf/

# Run the binary program produced by `go install`
#ENTRYPOINT ["./bin/mtd"]

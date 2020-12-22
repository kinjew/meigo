#!/bin/bash
if [ $# != 1 ] ; then
echo "USAGE: sh build.sh dev, stage or pub"
exit 1
fi;

PWD="$(pwd)";
#当前目录(scripts)
cd "$(dirname "$0")";

#环境变量
CHANNEL=$1
if [ $CHANNEL = "dev" ]; then
	export GIN_MODE=debug
fi;
if [ $CHANNEL = "stage" ]; then
	export GIN_MODE=test
fi;
if [ $CHANNEL = "pub" ]; then
	export GIN_MODE=release
fi;
echo "当前的环境是:$CHANNEL"

#在项目根目录下编译
cd ../
go build  -o bin/meigo main.go
go build  -o bin/cmd cmd/main.go
go build  -o bin/mtd daemon/mtd/main.go
echo '本地编译完成';
#交叉编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o linux_bin/meigo main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o linux_bin/cmd cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o linux_bin/mtd daemon/mtd/main.go
echo 'linux版本编译完成';

#配文文件、启停脚本、运行时目录等拷贝
rm -rf ./release
echo "移除旧的release目录完成"
mkdir release
echo "创建新的release目录完成"
cp -R ./bin ./release
cp -R linux_bin release
cp -R ./conf ./release
cp -R ./scripts ./release

echo '依赖拷贝完成';

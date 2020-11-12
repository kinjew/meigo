#! /bin/bash

SERVER="meigo"
BASE_DIR=$PWD
INTERVAL=1

# 命令行参数，需要手动指定
ARGS="-cDir conf/"

# shellcheck disable=SC2120
function build()
{
  PWD="$(pwd)";
  #当前目录(scripts)
  cd "$(dirname "$0")";

  #在项目根目录下编译
  cd ../
  go build  -o bin/meigo main.go
  go build  -o bin/cmd cmd/main.go
  echo '编译完成';

  #配文文件、启停脚本、运行时目录等拷贝
  cp -R bin release
  cp -R conf release
  cp -R scripts release

  echo '依赖拷贝完成';
}

function start()
{
  PID=$(lsof -i:8000|grep meigo|awk '{print $2}')

	if [[ $PID != "" ]]; then
		echo "$SERVER already running"
		lsof -i:8000|grep meigo|awk '{print "kill -USR2 "$2}'|bash
	fi

	$BASE_DIR/bin/$SERVER $ARGS &

	sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

function status()
{
  PID=$(lsof -i:8000|grep meigo|awk '{print $2}')
	if [[ $PID != "" ]]; then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	PID=$(lsof -i:8000|grep meigo|awk '{print $2}')
	if [[ $PID != "" ]]; then
		lsof -i:8000|grep meigo|awk '{print "kill -9 "$2}'|bash
	fi

	sleep $INTERVAL

	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi
}

case "$1" in
	'build')
	build
	;;
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	stop && start
	;;
	*)
	echo "usage: $0 {build|start|stop|restart|status}"
	exit 1
	;;
esac
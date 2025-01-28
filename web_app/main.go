package main

import (
	"fmt"
	"start/web_app/logger"
	"start/web_app/settings"
)

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 3.初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("load mysql failed, err:%v\n", err)
		return
	}
	// 4.初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("load redis failed, err:%v\n", err)
		return
	}

	// 除夕快乐
}

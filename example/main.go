package main

import (
	"flag"
	"fmt"
	"github.com/sestack/grf/example/global"
	"github.com/sestack/grf/example/initialize"
	"os"
)

func main() {
	var (
		err error
	)

	cfg := flag.String("c", "./config.yaml", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		fmt.Printf("version %s\n", global.Version)
		os.Exit(0)
	}

	// 初始化Viper
	global.VP, err = initialize.InitConfig(*cfg)
	if err != nil {
		fmt.Printf("解析配置文件失败:%v\v", err)
		os.Exit(1)
	}

	// gorm连接数据库
	global.DB, err = initialize.InitDB()
	if err != nil {
		fmt.Printf("初始化数据库失败：%v\v", err)
		os.Exit(1)
	}

	if global.DB != nil {
		// 初始化表
		if err = initialize.MigrateTables(global.DB); err != nil {
			fmt.Printf("初始化数据表失败：%v\v", err)
			os.Exit(1)
		}
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	httpServer := initialize.InitHttpServer()
	httpServer.Run(global.CONFIG.Http.Address)
}

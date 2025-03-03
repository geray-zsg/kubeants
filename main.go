package main

import (
	"kubeants.com/config"
	"kubeants.com/initiallize"
)

func main() {
	r := initiallize.Routers()
	// 初始化参数
	initiallize.Viper()
	initiallize.K8S()

	r.Run(config.CONF.System.Port)
}

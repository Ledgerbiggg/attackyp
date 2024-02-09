package main

import (
	"dos/config"
	"dos/job"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println("初始化成功")
	marshal, _ := json.Marshal(config.Configs)
	fmt.Println("配置文件：", string(marshal))

}

func main() {
	fmt.Println("开始攻击")
	job.AttackJob()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case <-interrupt:
		fmt.Println("结束攻击")
	}
}

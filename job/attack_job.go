package job

import (
	"dos/attack"
	"dos/config"
	"fmt"
)

func AttackJob() {

	// 创建一个新的 cron 实例
	c := NewWithSeconds()

	// 添加定时任务
	_, err := c.AddFunc(config.Configs.AttackCycle, func() {
		for i := 0; i < config.Configs.LoopCount; i++ {
			attack.Start()
		}
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// 启动 cron 服务
	c.Start()

	// 阻塞主线程
	select {}
}

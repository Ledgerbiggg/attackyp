/*
@author: ledger
@since: 2024/1/29
*/

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

// Configs 使用全局的配置变量
var Configs *YUPIAttackConfig

type YUPIAttackConfig struct {
	//  编程导航面试并发数量
	InterviewAttack int
	// 编程导航简历并发数量
	BiographicalNotesAttack int
	// 老鱼简历图床并发
	DrawingBedAttack int
	// 攻击周期
	AttackCycle string
	// 循环次数
	LoopCount int
}

// LoadConfig viper读取yaml
func LoadConfig() error {
	// yaml
	vconfig := viper.New()
	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 添加读取的配置文件路径
	vconfig.AddConfigPath(".")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	var cng YUPIAttackConfig
	if err := vconfig.Unmarshal(&cng); err != nil {
		log.Panicln("unmarshal cng file fail " + err.Error())
	}
	// 赋值全局变量
	Configs = &cng

	return err
}

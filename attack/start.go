package attack

import (
	"dos/config"
	"fmt"
)

var (
	InterviewAttackCount         int
	BiographicalNotesAttackCount int
	DrawingBedAttackCount        int
)

func Start() {

	// 编程导航攻击
	bcdhAttack := NewBCDHAttack()
	bcdhAttack.InterviewAttack(config.Configs.InterviewAttack)
	bcdhAttack.BiographicalNotesAttack(config.Configs.BiographicalNotesAttack)
	InterviewAttackCount += config.Configs.InterviewAttack
	BiographicalNotesAttackCount += config.Configs.BiographicalNotesAttack

	// 老鱼简历攻击
	lyjlAttack := NewLYJLAttack()
	lyjlAttack.DrawingBedAttack(config.Configs.DrawingBedAttack)
	DrawingBedAttackCount += config.Configs.DrawingBedAttack

	fmt.Println("InterviewAttackCount:", InterviewAttackCount)
	fmt.Println("BiographicalNotesAttackCount:", BiographicalNotesAttackCount)
	fmt.Println("DrawingBedAttack:", BiographicalNotesAttackCount)

}

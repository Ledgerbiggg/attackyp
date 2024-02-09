package utils

import (
	"dos/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestHttpDos_Get(t *testing.T) {
	get, err := NewHttpDos("https://www.baidu.com", nil, nil).Get()

	if err != nil {
		panic(err)
	}

	t.Log(string(get))
}
func TestHttpDos_Get1(t *testing.T) {
	headers := map[string]string{}
	headers["Content-Type"] = "application/image"
	get, err := NewHttpDos(
		"https://yupi-picture-1256524210.cos.ap-shanghai.myqcloud.com/640-20230912144925212.png",
		nil,
		headers).
		Get()
	if err != nil {
		panic(err)
	}
	fmt.Println("Byte Array:", string(get))
}

func TestHttpDos_Post2(t *testing.T) {
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
	post, err := NewHttpDos(""+
		"https://www.code-nav.cn/api/post/search/page/vo",
		marshal,
		headers).
		Post()
	if err != nil {
		t.Errorf("post error: %v", err)
	}

	fmt.Println("Byte Array:", string(post))
}

func TestHttpDos_Post1(t *testing.T) {
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
	post, err := NewHttpDos(""+
		"https://www.code-nav.cn/api/post/search/page/vo",
		marshal,
		headers).
		Post()
	if err != nil {
		t.Errorf("post error: %v", err)
	}

	fmt.Println("Byte Array:", string(post))
}

//
//func TestH(t *testing.T) {
//
//	url := "https://www.code-nav.cn/api/post/search/page/vo"
//	method := "POST"
//	req := model.NewInterviewReq(
//		"面试题",
//		10,
//		20,
//		1,
//		"createTime",
//		"descend",
//		model.Tags{Tags: []string{"Java"}})
//	headers := map[string]string{}
//	headers["Content-Type"] = "application/json"
//	marshal, _ := json.Marshal(req)
//
//	payload := strings.NewReader(string(marshal))
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
//	req.Header.Add("Content-Type", "application/json")
//	req.Header.Add("Accept", "*/*")
//	req.Header.Add("Host", "www.code-nav.cn")
//	req.Header.Add("Connection", "keep-alive")
//	req.Header.Add("Cookie", "SESSION=YWJmMjI0OTctZjc4YS00YzgzLWI1NjItMDA4OGM4ZTljN2Ni")
//
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//
//}

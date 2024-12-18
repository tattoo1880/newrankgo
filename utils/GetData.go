package utils

import (
	"encoding/json"
	"fmt"
	"gotest/model"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RequestData struct {
	Filename string
	URL      string
}

type JsonData struct {
	RankType  int    `json:"rankType"`
	RankDate  string `json:"rankDate"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	Start     int    `json:"start"`
	PhotoType string `json:"photoType"`
}

type JsonResponse struct {
	Data struct {
		List []map[string]interface{} `json:"list"`
	} `json:"data"`
}

func GetData(wg *sync.WaitGroup, dict RequestData) {
	defer wg.Done()

	// 请求头
	headers := map[string]string{
		"Accept":             "*/*",
		"Accept-Language":    "zh-CN,zh;q=0.9",
		"Connection":         "keep-alive",
		"Content-Type":       "application/json;charset=UTF-8",
		"n-token":            "3b2f8f99af0545cc989cfae76477d9bf",
		"Origin":             "https://www.newrank.cn",
		"Referer":            "https://www.newrank.cn/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		"sec-ch-ua":          `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"macOS"`,
	}

	// 请求体
	//jsonData := `{
	//	"rankType": 1,
	//	"rankDate": "2024-12-17",
	//	"type": "",
	//	"size": 50,
	//	"start": 1,
	//	"photoType": ""
	//}`

	var jsonData JsonData
	jsonData.RankType = 1
	//今天的年月日
	jsonData.RankDate = time.Now().Format("2006-01-02")
	jsonData.Type = ""
	jsonData.Size = 50
	jsonData.Start = 1
	jsonData.PhotoType = ""

	jsonBytes, _ := json.Marshal(jsonData)

	json_string := string(jsonBytes)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", dict.URL, strings.NewReader(json_string))
	if err != nil {
		fmt.Printf("Error creating request for %s: %v\n", dict.Filename, err)
		return
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to %s: %v\n", dict.Filename, err)
		return
	}
	defer resp.Body.Close()

	// 解析响应 JSON
	var responseJson JsonResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseJson); err != nil {
		fmt.Printf("Error decoding response for %s: %v\n", dict.Filename, err)
		return
	}

	// 打印解析后的数据
	fmt.Printf("Data for %s:\n", dict.Filename)
	if dict.Filename == "xhs" {
		for _, item := range responseJson.Data.List {

			var xhs model.Xhs

			jsonBytes, _ := json.Marshal(item)
			json.Unmarshal(jsonBytes, &xhs)
			fmt.Println(xhs)
			//todo 保存到数据库
			err := xhs.Create()
			if err != nil {
				return
			}

			//todo break
			//break
		}
	} else if dict.Filename == "weixin" {
		for _, item := range responseJson.Data.List {
			jsonBytes, _ := json.Marshal(item)
			var weixin model.Weixin
			json.Unmarshal(jsonBytes, &weixin)
			fmt.Println(weixin)
			//todo 保存到数据库
			err := weixin.Create()
			if err != nil {
				return
			}
			//break
		}
	} else if dict.Filename == "dy" {
		for _, item := range responseJson.Data.List {
			for k, v := range item {
				fmt.Printf("%s: %v\n", k, v)
			}

			var dy model.Dy
			jsonBytes, _ := json.Marshal(item)
			if err := json.Unmarshal(jsonBytes, &dy); err != nil {
				fmt.Println("error:", err)
				return
			}
			err := dy.Create()
			if err != nil {
				return
			}
			//break
		}
	} else if dict.Filename == "sph" {
		for _, item := range responseJson.Data.List {
			//for k, v := range item {
			//	fmt.Printf("%s: %v\n", k, v)
			//}
			var sph model.Sph
			jsonBytes, _ := json.Marshal(item)
			if err := json.Unmarshal(jsonBytes, &sph); err != nil {
				fmt.Println("error:", err)
				return
			}
			err := sph.Create()
			if err != nil {
				return
			}
		}
	} else if dict.Filename == "ks" {
		for _, item := range responseJson.Data.List {
			for k, v := range item {

				fmt.Printf("%s: %v\n", k, v)
			}

			var ks model.Ks
			jsonBytes, _ := json.Marshal(item)
			if err := json.Unmarshal(jsonBytes, &ks); err != nil {
				fmt.Println("error:", err)
				return
			}
			err := ks.Create()
			if err != nil {
				return
			}

		}
	} else if dict.Filename == "bilibili" {
		for _, item := range responseJson.Data.List {
			//for k, v := range item {
			//	fmt.Printf("%s: %v\n", k, v)
			//}
			var bilibili model.Bil
			jsonBytes, _ := json.Marshal(item)
			if err := json.Unmarshal(jsonBytes, &bilibili); err != nil {
				fmt.Println("error:", err)
				return
			}
			err := bilibili.Create()
			if err != nil {
				return
			}
		}
	}
}

func Doit() {
	// 定义 URL 列表
	urlList := []RequestData{
		{"xhs", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getXhsHotContent"},
		{"weixin", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getWxHotContent"},
		{"dy", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getDyHotContent"},
		{"sph", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getSphHotContent"},
		{"ks", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getKsHotContent"},
		{"bilibili", "https://gw.newrank.cn/api/mainRank/nr/mainRank/hotContent/getBiliHotContent"},
	}

	// 使用 sync.WaitGroup 并发请求
	var wg sync.WaitGroup
	for _, data := range urlList {
		wg.Add(1)
		go GetData(&wg, data)
	}

	wg.Wait()
	fmt.Println("All requests finished!")
}

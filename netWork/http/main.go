package main

import (
	"fmt"
	"io"
	"net/http"
)

// 使用get
func getM(url string) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("请求获取失败", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("请求获取失败", err)
	}

	fmt.Println("请求结果：", string(data))
}

// 使用原始方法，先构造一个request ，然后用DO()发送
func originRequest(url string, method string) {
	//
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("请求获取失败", err)
	}

	http.DefaultClient.Do(request)

}
func main() {

}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// path, err := os.Getwd()
	// fmt.Print(path)
	//使用os包的ReadFile可以一次性读出文件内容
	data, err := os.ReadFile("../../fileWrite/WritTest.txt")
	err2 := os.WriteFile("../../fileWrite/WritTest.txt", []byte("\n其实你不是不爱了吧，是有好多事没处理，\n怎么你闭口不语，是不是我正好说中你的心，就承认还是在意吧，就偏偏我也可以"), 0666)

	if err != nil || err2 != nil {
		log.Fatal("打开文件失败", err)
	}

	fmt.Println(string(data))
}

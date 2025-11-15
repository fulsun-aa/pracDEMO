package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// read
// 先用bufio.NewReader(os.File)创建一个Reader
// ReadString(”)是使用分隔符读取每次读到分隔符就会返回一次读取的数据。
func rd(f *os.File) {
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		//这里采用分隔符读取办法，会把分隔符也读进去
		line, err := reader.ReadString('\n')
		//要判断读取错误不是因为读到结尾导致的
		if err != nil && err != io.EOF {
			log.Fatal("读取失败：,", err)
			return
		}
		fmt.Print(line)
		if err == io.EOF {
			fmt.Println("\n读取结束")
			return
		}
	}
}

func main() {
	file, err := os.Open("../fileWrite/WritTest.txt")
	if err != nil {
		log.Fatal("打开文件失败,", err)
		return
	}

	rd(file)
}

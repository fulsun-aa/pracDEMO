package main

import (
	"fmt"

	"log"
	"os"
)

func main() {
	dir, err := os.ReadDir("./")
	if err != nil {
		log.Fatal("读取目录失败,", err)
		return
	}
	for _, entry := range dir {
		fmt.Println("文件名:", entry.Name())
		fmt.Println("文件类型:", entry.Type())
		fmt.Println("------------------")
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("safsafadsgsdagasdgsadfwfweafgawg")
	printNext5Bytes(r)
	fmt.Println("让游标回退三个字节")
	for i := 0; i < 3; i++ {
		r.UnreadByte()
	}
	printNext5Bytes(r)

	fmt.Println("开始读取中文")

	r2 := strings.NewReader("一二三四五六七八九十")
	printNext5chars(r2)
	fmt.Println("让游标回退三个字节")
	for i := 0; i < 3*3; i++ {
		r2.UnreadRune()
	}
	printNext5chars(r2)

}

// reader.ReadRune()可以读取UTF-8
func printNext5chars(reader *strings.Reader) {
	for i := 0; i < 5; i++ {

		charsData, _, _ := reader.ReadRune()
		fmt.Println(string(charsData))
	}
}

// reader.ReadByte（）不能读取UTF-8
func printNext5Bytes(reader *strings.Reader) {
	for i := 0; i < 5; i++ {
		//strings.Reader 底层是用的切片，将原始字符内容复制到字节切片中去。所以如果原始字符是中文，则读取会产生乱码
		//strings.Reader 是根据游标来读取的
		byteData, _ := reader.ReadByte()
		fmt.Println(string(byteData))
	}
}

package main

import (
	"log"
	"os"
)

// 普通读取的方法
func generalRead(f *os.File) {
	defer f.Close()
	bytes := make([]byte, 1024)
	//Read方法会返回读取的字节数，和一个错误信息,读取的内容会存储在bytes切片中
	n, err := f.Read(bytes)

	if err == nil {
		log.Printf("共读取%d个字符，文件内容是：\n%s", n, string(bytes))
	} else {
		log.Printf("读取文件失败: %v", err)
	}
}

// 循环读取,由于字节切片大小是无法确定的，所以我们可以用循环读取，每次对返回值err判断，读取错误则跳出循环
func circleRead(f *os.File) {
	defer f.Close()
	var content []byte

	for {
		bytes := make([]byte, 10)
		n, err := f.Read(bytes)
		//在 Go 语言中，... 在这里是切片展开运算符，作用是将一个切片的所有元素 “打散”，作为独立的参数传递给可变参数函数（如 append）

		if err == nil {
			content = append(content, bytes[:n]...)
			log.Printf("共读取%d个字符，文件内容是：\n%s", n, string(bytes))
		} else {

			log.Printf("读取文件失败: %v", err)
			break
		}
	}
	//Read方法会返回读取的字节数，和一个错误信息,读取的内容会存储在bytes切片中
	log.Printf("最终读取到的文件内容是：\n%s", string(content))
}

func main() {
	//会返回一个os.File类型的指针，并且os.File实现了io.Reader接口，所以可以使用Read方法读取文件内容
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	//general reading
	//generalRead(f)
	//circle reading
	circleRead(f)
}

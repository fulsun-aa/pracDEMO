package main

import (
	"log"
	"os"
)

func generalWrite(f *os.File, str string) {
	defer f.Close()

	n, err := f.Write([]byte(str))

	if err == nil {
		log.Printf("共写入了%d个字节，内容是：\n%s", n, str)
	} else {
		log.Fatal("写入文件失败,", err)
	}
}

func tailWrite(path string, str string) {
	//os.O_RDWR|os.O_CREATE|os.O_APPEND 是以读写方式|文件不存在就创建方式|追加方式，0666指的是文件的权限属性
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("创建文件失败,", path, err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte(str))

	if err == nil {
		log.Printf("追加写入了%d个字节，内容是：\n%s", n, str)
	} else {
		log.Fatal("写入文件失败,", err)
	}
}

func main() {
	path := "WritTest.txt"
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("创建文件失败,", err)
		return
	}

	generalWrite(file, "baby 我们的爱情好像跳楼机，让我突然的升空又急速落地~~")
	tailWrite(path, "你带给我一束风光，劫后余生好难呼吸，那天的天气难得放晴，你说的话却把我困在雨季")

}

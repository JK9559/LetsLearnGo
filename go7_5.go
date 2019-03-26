package main

import (
	"fmt"
	"os"
)

// 目录操作
func dirOp() {
	// 创建目录 设置权限为 0777
	os.Mkdir("astaxie", 0777)
	// 创建多级目录 设置权限为0777
	os.MkdirAll("astaxie/test1/test2", 0777)
	// 删除单级目录 如果目录下有文件或者其他目录会报错
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	// 删除多级子目录 如果为单个名称 那么该目录下的子目录全被删除
	os.RemoveAll("astaxie")
}

func writeFl() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		// 写入string信息到文件
		fout.WriteString("Just a test!\r\n")
		// 写入byte类型的信息到文件
		fout.Write([]byte("Just a byteTest!\r\n"))
	}
}

func readFl() {
	userFile := "astaxie.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func go7_5() {
	readFl()
}

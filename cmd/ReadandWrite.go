package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// ReadFile 读取文件列表返回数组
func ReadFile(File string) string {
	f, err := os.Open(File)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

// Write 写文件
func Write(output string, number int) {
	//创建文件夹
	_ = os.Mkdir("./results", os.ModePerm)
	var filename string
	if number >= 1 {
		filename = "./results/" + strconv.Itoa(number) + ".txt"
	} else if number == 0 {
		filename = "./results/" + "log.txt"
	} else {
		filename = "./results/" + "enbase64.txt"
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("文件错误,错误为:%v\n", err)
		return
	}
	str := []byte(output)
	_, _ = file.Write(str) //将str字符串的内容写到文件中，强制转换为byte，因为Write接收的是byte。
}

package cmd

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"path"
	"strconv"
	"strings"
)

func Run() {
	if File != "" {
		//判断文件后缀类型是不是txt
		if strings.Contains(Filename(File), ".txt") == true {
			Partition()
		} else {
			Partitions(ReadFile(File), Number)
		}
	} else if EncodeFile != "" {
		EnBase64()
		color.HiGreen.Println("[+] 编码文件成功，请见results目录！")
		color.HiGreen.Println("上传base64编码后的文件可使用还原命令：" +
			"certutil -decode test.txt test.exe\n")

	} else {
		fmt.Println("啥也没有,干点啥呢！！")
	}

}

// Partitions  分割木马文件并生成写入命令
func Partitions(str string, Number int) {
	//替换 " 为 '
	newstrs := strings.Replace(str, `"`, `'`, -1)
	wins := SplitSubN(newstrs, Number)
	fmt.Println("windows写入多个文件后合并命令：")
	Write("windows写入多个文件后合并命令：\n", 0)

	var winstr []string
	for i := 0; i < len(wins); i++ {
		req := "echo|set /p=\"" + wins[i] + "\">" + strconv.Itoa(i) + ".txt"
		winstr = append(winstr, req)
	}
	for _, v := range winstr {
		fmt.Println(v)
		Write(v+"\n", 0)
	}
	fmt.Println("使用copy命令合并文件！")
	color.HiGreen.Println("示例：copy 0.txt + 1.txt out.txt！")
	Write("使用copy命令合并文件！\n", 0)
	Write("示例：copy 0.txt + 1.txt out.txt！\n", 0)
	fmt.Println()
	fmt.Println("windows追加字符：")
	Write("windows追加字符：\n", 0)
	var winstrs []string
	for i := 0; i < len(wins); i++ {
		req := "echo|set /p=\"" + wins[i] + "\">>" + "test" + Filename(File)
		winstrs = append(winstrs, req)
	}
	for _, v := range winstrs {
		fmt.Println(v)
		Write(v+"\n", 0)
	}
	fmt.Println("================================================================")
	Write("================================================================\n", 0)

	//________________________________________________________________

	newlinuxstrs := strings.Replace(str, `'`, `"`, -1)
	lins := SplitSubN(newlinuxstrs, Number)
	fmt.Println()
	fmt.Println("linux追加写入命令：")
	Write("linux追加写入命令：\n", 0)
	var linstr []string
	for i := 0; i < len(lins); i++ {
		req := "echo -n '" + lins[i] + "' >>" + "1.txt"
		linstr = append(linstr, req)
	}
	for _, v := range linstr {
		fmt.Println(v)
		Write(v+"\n", 0)
	}

	fmt.Printf("\nlinux写入多个文件后合并：\n")
	Write("linux写入多个文件后合并：\n", 0)
	linsr := SplitSubN(str, Number)
	fmt.Println("linux分割写入后合并命令：")
	Write("linux分割写入后合并命令：\n", 0)
	var linstrs []string
	for i := 0; i < len(linsr); i++ {
		//fmt.Println("echo '"+ lins[i] +"' >" + strconv.Itoa(i) + ".txt")
		req := "echo -n '" + linsr[i] + "' >" + strconv.Itoa(i) + ".txt"
		linstrs = append(linstrs, req)
	}
	for _, v := range linstrs {
		fmt.Println(v)
		Write(v+"\n", 0)
	}
	//fmt.Printf("示例命令：paste -d '' 1.txt 2.txt > 3.txt\n")
	color.HiGreen.Printf("合并文件示例命令：paste -d '' 1.txt 2.txt > 3.txt\n")
	Write("合并文件示例命令：paste -d '' 1.txt 2.txt > 3.txt\n", 0)

}

// Partition 分割txt文件
func Partition() {
	str := SplitSubN(ReadFile(File), Number)
	fmt.Printf("分割为 %v 个\n", len(str))
	for i, v := range str {
		fmt.Println(v)
		Write(v, i)
	}
}

// SplitSubN 按照长度切割字符串
func SplitSubN(s string, n int) []string {
	var sub string
	var subs []string

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}
	return subs
}

// Filename 读取文件后缀
func Filename(File string) string {

	filenameWithSuffix := path.Base(File)
	fileSuffix := path.Ext(filenameWithSuffix)

	return fileSuffix
}

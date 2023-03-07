package cmd

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"path"
	"strconv"
	"strings"
)


func Run(){
	if File != "" {
		if strings.Contains(Filename(File),".txt") ==true {
			Partition()
		}else {
			Partitions(ReadFile(File),Number)
		}
	}else {
		fmt.Println("啥也没有,干点啥呢！！")
	}



}

// Partitions  分割并生成写入命令
func Partitions(str string,Number int){
	newstrs := strings.Replace(str,`"`,`'`,-1)
	wins :=SplitSubN(newstrs,Number)
	//strs := strings.Replace(str,"<","^<",-1)
	//newstrs := strings.Replace(strs,">","^>",-1)
	//wins := SplitSubN(newstrs,10)
	//for _,v := range wins {
	//	fmt.Println(v)
	//}
	//fmt.Printf("\n分成了%v个片段\n\n", len(wins))
	fmt.Println("windows写入多个文件后合并命令：")

	var winstr []string
	for i := 0; i< len(wins); i++ {
		//fmt.Println("echo "+ wins[i] +" >" + strconv.Itoa(i) + ".txt")
		req := "echo|set /p=\""+wins[i] + "\">" + strconv.Itoa(i) + ".txt"
		winstr = append(winstr,req)
	}
	for _,v := range winstr {
		fmt.Println(v)
	}
	fmt.Println("使用copy命令合并文件！")
	//fmt.Println("示例：copy 0.txt + 1.txt out.txt！")
	color.HiGreen.Println("示例：copy 0.txt + 1.txt out.txt！")
	fmt.Println()
	fmt.Println("windows追加字符：")
	var winstrs []string
	for i := 0; i< len(wins); i++ {
		//fmt.Println("echo "+ wins[i] +" >" + strconv.Itoa(i) + ".txt")
		req := "echo|set /p=\""+wins[i] + "\">>" + "test"+ Filename(File)
		winstrs = append(winstrs,req)
	}
	for _,v := range winstrs {
		fmt.Println(v)
	}
	fmt.Println("================================================================")

	//________________________________________________________________

	newlinuxstrs := strings.Replace(str,`'`,`"`,-1)
	lins := SplitSubN(newlinuxstrs,Number)
	//for _,v := range lins {
	//	fmt.Println(v)
	//}
	//fmt.Printf("分成了%v个片段\n\n", len(lins))
	fmt.Println()
	fmt.Println("linux追加写入命令：")
	var linstr []string
	for i := 0; i< len(lins); i++ {
		//fmt.Println("echo '"+ lins[i] +"' >" + strconv.Itoa(i) + ".txt")
		req := "echo -n '"+ lins[i] +"' >>" + "1.txt"
		linstr = append(linstr,req)
	}
	for _, v := range linstr{
		fmt.Println(v)
	}

	fmt.Printf("\nlinux写入多个文件后合并：\n")
	linsr := SplitSubN(str,Number)
	//for _,v := range lins {
	//	fmt.Println(v)
	//}
	//fmt.Printf("分成了%v个片段\n\n", len(lins))

	fmt.Println("linux分割写入后合并命令：")
	var linstrs []string
	for i := 0; i< len(linsr); i++ {
		//fmt.Println("echo '"+ lins[i] +"' >" + strconv.Itoa(i) + ".txt")
		req := "echo -n '"+ linsr[i] +"' >" + strconv.Itoa(i) + ".txt"
		linstrs = append(linstrs,req)
	}
	for _, v := range linstrs{
		fmt.Println(v)
	}
	//fmt.Printf("示例命令：paste -d '' 1.txt 2.txt > 3.txt\n")
	color.HiGreen.Printf("合并文件示例命令：paste -d '' 1.txt 2.txt > 3.txt\n")
}

// Partition 分割txt文件
func Partition() {
	str := SplitSubN(ReadFile(File),Number)
	fmt.Printf("分割为 %v 个\n",len(str))
	for i,v := range str {
		fmt.Println(v)
		Write(v,i)
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


func Filename(File string) string{

	filenameWithSuffix := path.Base(File)
	fileSuffix := path.Ext(filenameWithSuffix)

	return fileSuffix
}
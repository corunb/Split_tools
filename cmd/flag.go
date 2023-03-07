package cmd

import (
	"flag"
	"github.com/gookit/color"
)

var Number int
var File string



func init() {
	flag.IntVar(&Number,"n",32,"设置长度进行切割，默认32")
	flag.StringVar(&File, "f","","指定分割的文本")
	flag.Parse()


	logo := `
  ____        _ _ _      _              _     
 / ___| _ __ | (_) |_   | |_ ___   ___ | |___ 
 \___ \| '_ \| | | __|  | __/ _ \ / _ \| / __|
  ___) | |_) | | | |_   | || (_) | (_) | \__ \
 |____/| .__/|_|_|\__|___\__\___/ \___/|_|___/
       |_|          |_____|                   

[+] code by Corun V1.0
[+] https://github.com/corunb/Split_tools
`
	color.HiGreen.Println(logo)


}



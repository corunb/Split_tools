package cmd

import "encoding/base64"

func EnBase64() {
	en64 := ReadFile(EncodeFile)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString([]byte(en64))
	Write(sourcestring, -1)
}

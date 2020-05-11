package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"sort"
)

//全局变量
var port, token string = "4521", "123456789"

func main() {
	config := ReadProfile()
	port = config.Port
	token = config.Token
	// 绑定路由
	http.HandleFunc("/", checkout)
	fmt.Printf("在端口 %s 上启动服务器...\n", config.Port)
	// 启动监听=j
	err2 := http.ListenAndServe(":"+port, nil)
	if err2 != nil {
		fmt.Println("服务器启动失败！")
	}
}

//checkout 用于监听和处理绑定的接口上的数据
func checkout(response http.ResponseWriter, request *http.Request) {
	//解析URL参数
	err := request.ParseForm()
	if err != nil {
		fmt.Println("URL解析失败！")
		return
	}
	// 获取参数
	signature := request.FormValue("signature")
	timestamp := request.FormValue("timestamp")
	nonce := request.FormValue("nonce")
	echostr := request.FormValue("echostr")
	//将token、timestamp、nonce三个参数进行字典序排序
	var tempArray = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		_, err := response.Write([]byte(echostr))
		if err != nil {
			fmt.Println("响应失败。。。")
		}
	} else {
		fmt.Println("验证失败")
	}
	if request.Method == "POST" {
		text := parsingXMLtext(request)
		if text != nil {
			fmt.Printf("收到来自用户 %s 的消息：%s\n", text.FromUserName, text.Content)
			//responseTextBody 回复用户的消息
			responseTextBody, err := makeXMLtext(text.ToUserName, text.FromUserName, "喵, 吾乃FBK。我收到你的这个消息辣："+text.Content)
			if err != nil {
				log.Println("Wechat Service: makeXMLtext error: ", err)
				return
			}
			response.Header().Set("Content-Type", "text/xml")
			fmt.Fprintf(response, string(responseTextBody))
		}
	}
}

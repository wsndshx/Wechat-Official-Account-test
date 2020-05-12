package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//XMLbase 消息的基础内容,用于判断消息类型
type XMLbase struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	PicUrl       string
	Mediald      string
	MsgId        int64
}

//XMLtext 用户消息结构体(text)
/* type XMLtext struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int64
} */

//XMLimage 用户消息结构体(image)
/* type XMLimage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	PicUrl       string
	Mediald      string
	MsgId        int64
} */

//reXMLimage 用于回复消息的结构体(image)
type reXMLimage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   time.Duration
	MsgType      CDATA
	PicUrl       CDATA
	Mediald      CDATA
	MsgId        int64
}

//reXMLtext 用于回复消息的结构体(text)
type reXMLtext struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   time.Duration
	MsgType      CDATA
	Content      CDATA
	MsgId        int64
}

//CDATA 一个格式
type CDATA struct {
	Text string `xml:",innerxml"`
}

//valueCDATA 给消息加上CDATA
func valueCDATA(v string) CDATA {
	return CDATA{"<![CDATA[" + v + "]]>"}
}

//ParsingXMLtext 用于解析并返回格式化后的用户消息(text)
/* func parsingXMLtext(response *http.Request) *XMLtext {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &XMLtext{}
	xml.Unmarshal(body, requestBody)
	return requestBody
} */

//ParsingXMLbase 用于解析并返回接收的任何类型的消息
func parsingXMLbase(response *http.Request) *XMLbase {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &XMLbase{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

//parsingXMLimage
/* func parsingXMLimage(response *http.Request) *XMLimage {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &XMLimage{}
	xml.Unmarshal(body, requestBody)
	return requestBody
} */

//makeXMLtext 用于生成包含回复的内容的xml
func makeXMLtext(fromUserName, toUserName, content string) ([]byte, error) {
	XMLtext := &reXMLtext{}
	XMLtext.FromUserName = valueCDATA(fromUserName)
	XMLtext.ToUserName = valueCDATA(toUserName)
	XMLtext.MsgType = valueCDATA("text")
	XMLtext.Content = valueCDATA(content)
	XMLtext.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(XMLtext, " ", "  ")
}

//用于处理转发消息
func makeXML(fromUserName, toUserName, content string) ([]byte, error) {
	XMLtext := &reXMLtext{}
	XMLtext.FromUserName = valueCDATA(fromUserName)
	XMLtext.ToUserName = valueCDATA(toUserName)
	XMLtext.MsgType = valueCDATA("text")
	XMLtext.Content = valueCDATA(content)
	XMLtext.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(XMLtext, " ", "  ")
}

//replyXMLtext 消息回复函数(text)
//[接受到的消息XML] [http.ResponseWriter] [要回复的内容]
func replyXMLtext(text *XMLbase, response http.ResponseWriter, content string) {
	responseTextBody, err := makeXMLtext(text.ToUserName, text.FromUserName, content)
	if err != nil {
		log.Println("Wechat Service: makeXMLtext error: ", err)
		return
	}
	response.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(response, string(responseTextBody))
}

//forwardMessage 将消息转发至微信客服平台
func forwardMessage(text *XMLbase, response http.ResponseWriter) {
	responseTextBody, err := makeXMLtext(text.ToUserName, text.FromUserName, "transfer_customer_service")
	if err != nil {
		log.Println("Wechat Service: makeXMLtext error: ", err)
		return
	}
	response.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(response, string(responseTextBody))
}

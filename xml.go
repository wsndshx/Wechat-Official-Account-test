package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//XMLtext 用户消息模板(text)
type XMLtext struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type reXMLtext struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   time.Duration
	MsgType      CDATA
	Content      CDATA
	MsgId        int
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
func parsingXMLtext(response *http.Request) *XMLtext {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &XMLtext{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

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

//replyXMLtext 消息回复函数
//[接受到的消息XML] [http.ResponseWriter] [要回复的内容]
func replyXMLtext(text *XMLtext, response http.ResponseWriter, content string) {
	responseTextBody, err := makeXMLtext(text.ToUserName, text.FromUserName, content)
	if err != nil {
		log.Println("Wechat Service: makeXMLtext error: ", err)
		return
	}
	response.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(response, string(responseTextBody))
}

package main

import (
	"encoding/xml"
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

func makeXMLtext(fromUserName, toUserName, content string) ([]byte, error) {
	XMLtext := &XMLtext{}
	XMLtext.FromUserName = fromUserName
	XMLtext.ToUserName = toUserName
	XMLtext.MsgType = "text"
	XMLtext.Content = content
	XMLtext.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(XMLtext, " ", "  ")
}

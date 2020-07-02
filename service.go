package sfexpress

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Head string
type ErrCode string
type ErrMsg string

func (c Head) IsSuccess() bool {
	return c == HeadSuccess
}

type Err struct {
	ErrCode ErrCode `xml:"code,attr"`
	ErrMsg  ErrMsg `xml:",chardata"`
}

type ResponseXml struct {
	XMLName xml.Name `xml:"Response"`
	Head    Head   `xml:"Head"`
	Err     Err      `xml:"ERROR"`
	Body    Body     `xml:"Body"`
}

const (
	HeadSuccess                         Head    = "OK"   //交易成功
	HeadERR                             Head    = "ERR"  //系统或业务异常,交易失败
	ErrCodeAddrRequired                 ErrCode = "010"  //寄件地址不能为空
	ErrCodeContractNameRequired         ErrCode = "1011" //寄件联系人不能为空
	ErrCodeTelPhoneRequired             ErrCode = "1012" //寄件电话不能为空
	ErrCodeReceiverAddrRequired         ErrCode = "1014" //到件地址不能为空
	ErrCodeReceiverContractNameRequired ErrCode = "1015" //到件联系人不能为空
	ErrCodeReceiverTelPhoneRequired     ErrCode = "1016" //到件联系人不能为空
	ErrCodeCargoRequired                ErrCode = "1017" //到件联系人不能为空
)

func (c *Config) postData(xml []byte) string {
	var buf bytes.Buffer
	buf.Write(xml)
	buf.Write([]byte(c.CheckWord))
	s := c.sign(buf.Bytes())
	v := url.Values{}
	v.Add("xml", fmt.Sprintf("%s", xml))
	v.Add("verifyCode", s)
	return v.Encode()
}

func (c *Config) sign(s []byte) string {
	has := md5.Sum(s)
	md5str := fmt.Sprintf("%s", has)
	encodeString := base64.StdEncoding.EncodeToString([]byte(md5str))
	return encodeString
}

func (c *Config) doRequest(requestBody RequestBody, serviceName ServiceName) (*Body, error) {
	reqXml := RequestXml{
		XMLName:     xml.Name{Space: "Request"},
		Service:     serviceName,
		Lang:        RequestServiceLang,
		Head:        RequestServiceHead,
		RequestBody: requestBody,
	}
	xmlByte, err := xml.Marshal(reqXml)
	if err != nil {
		return nil, err
	}

	p := c.postData(xmlByte)
	body := strings.NewReader(p)
	resp, err := http.Post(ServiceURL, "application/x-www-form-urlencoded;charset=utf-8", body)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	r := new(ResponseXml)
	err = xml.Unmarshal([]byte(b), r)
	if err != nil {
		return nil, err
	}
	if r.Head != "OK" {
		return nil, r.Err
	}

	return &r.Body, nil
}

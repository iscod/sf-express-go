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

func (c *Config) doRequest(requestBody RequestBody) (*Body, error) {
	reqXml := RequestXml{
		XMLName:     xml.Name{Space: "Request"},
		Service:     OrderServiceName,
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

package sfexpress

import (
	"bytes"
	"encoding/xml"
	"crypto/md5"
	"encoding/base64"
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

	s := c.sgin(buf.Bytes())
	v := url.Values{}
	v.Add("xml", fmt.Sprintf("%s", xml))
	v.Add("verifyCode", s)
	return v.Encode()
}

func (c *Config) sgin(s []byte) string {
	has := md5.Sum(s)
	md5str := fmt.Sprintf("%s", has)
	encodeString := base64.StdEncoding.EncodeToString([]byte(md5str))
	return encodeString
}

func (c *Config) OrderPush(orderid string, jUser UserInfo, dUser UserInfo, expressType int, payMethod int, addedService AddedService, cargo string) (*OrderResponse, error) {
	req_xml := RequestXml{
		XMLName: xml.Name{Space: "Request"},
		Service: OrderServiceName,
		Lang:    RequestServiceLang,
		Head:    RequestServiceHead,
		RequestBody: RequestBody{
			OrderPush: OrderPush{
				OrderId:        orderid,
				ExpressType:    expressType,
				PayMethod:      payMethod,
				JCompany:       jUser.Company,
				JContact:       jUser.Contact,
				JTel:           jUser.Tel,
				JMobile:        jUser.Mobile,
				JProvince:      jUser.Province,
				JCity:          jUser.City,
				JCounty:        jUser.County,
				JAddress:       jUser.Address,
				DCompany:       dUser.Company,
				DContact:       dUser.Contact,
				DTel:           dUser.Tel,
				DMobile:        dUser.Mobile,
				DProvince:      dUser.Province,
				DCity:          dUser.City,
				DCounty:        dUser.County,
				DAddress:       dUser.Address,
				ParcelQuantity: "1",
				Custid:         c.Custid,
				AddedService:   addedService,
				Cargo:          "sss",
			},
		},
	}
	xml_byte, err := xml.Marshal(req_xml)
	if err != nil {
		return nil, err
	}

	p := c.postData(xml_byte)
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

	return &r.Body.OrderResponse, nil
}

func (c *Config) OrderQuery(orderid string) (*OrderResponse, error) {
	req_xml := RequestXml{
		XMLName:     xml.Name{Space: "Request"},
		Service:     OrderSearchServiceName,
		Lang:        "zh-CN",
		Head:        "SLKJ2019",
		RequestBody: RequestBody{OrderSearch: OrderSearch{OrderId: orderid}},
	}
	xml_byte, err := xml.Marshal(req_xml)
	if err != nil {
		return nil, err
	}
	p := c.postData(xml_byte)
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

	return &r.Body.OrderResponse, nil
}

func (c *Config) OrderConfirm(orderid string, dealtype string) (*OrderResponse, error) {
	req_xml := RequestXml{
		XMLName: xml.Name{Space: "Request"},
		Service: "OrderConfirmService",
		Lang:    "zh-CN",
		Head:    "SLKJ2019",
		RequestBody: RequestBody{
			OrderConfirm: OrderConfirm{
				OrderId:  orderid,
				DealType: dealtype,
			},
		},
	}
	xml_byte, err := xml.Marshal(req_xml)
	if err != nil {
		return nil, err
	}

	p := c.postData(xml_byte)
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

	return &r.Body.OrderResponse, nil
}

func (c *Config) OrderRouteService(no string, trackingType string, methodType string) (*RouteResponse, error) {
	req_xml := RequestXml{
		XMLName:     xml.Name{Space: "Request"},
		Service:     "RouteService",
		Lang:        "zh-CN",
		Head:        "SLKJ2019",
		RequestBody: RequestBody{RouteRequest: RouteRequest{TrackingNumber: no, TrackingType: trackingType, MethodType: methodType}},
	}
	xml_byte, err := xml.Marshal(req_xml)
	if err != nil {
		return nil, err
	}
	p := c.postData(xml_byte)
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

	return &r.Body.RouteResponse, nil
}

func (c *Config) OrderRouteQueryByOrderNo(orderno string) (*RouteResponse, error) {
	o, err := c.OrderRouteService(orderno, "2", "1")
	return o, err
}

func (c *Config) OrderRouteQueryByMailNo(mailno string) (*RouteResponse, error) {
	o, err := c.OrderRouteService(mailno, "1", "1")
	return o, err
}

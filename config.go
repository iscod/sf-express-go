package sfexpress

import (
	"fmt"
	"encoding/xml"
)

var (
	OrderServiceName       = "OrderService"
	OrderSearchServiceName = "OrderSearchService"
	RequestServiceLang     = "zh-CN"
	RequestServiceHead     = "SLKJ2019"
	ServiceURL             = "http://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
	ServiceURLHttps        = "https://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
)

type ExpressTypeCode struct {
	Id   int
	Code string
}

var ExpressType = map[int]ExpressTypeCode{
	1:  {Id: 1, Code: "T4"},   //顺丰标快
	2:  {Id: 2, Code: "T6"},   //顺丰标快（陆运）
	5:  {Id: 5, Code: "T8"},   //顺丰次晨
	6:  {Id: 6, Code: "T1"},   //顺丰即日
	9:  {Id: 9, Code: "T13"},  //顺丰国际小包（平邮)
	10: {Id: 10, Code: "T14"}, //顺丰国际小包（挂号)
	12: {Id: 12, Code: "T4"},  //医药专递
	13: {Id: 13, Code: "T6"},  //物流普运
}

type Config struct {
	ClientCode string `NTL`
	CheckWord  string `P9CUF8XuqBUH2ZaFpwjyqTusx12zr4uM`
	Custid     string
	Lang       string `zh-CN`
}

type export struct {
	name string
	url  string
}

type Err struct {
	ErrMsg  string `xml:",chardata"`
	ErrCode string `xml:"code,attr"`
}

type ResponseXml struct {
	XMLName xml.Name `xml:"Response"`
	Head    string   `xml:"Head"`
	Err     Err      `xml:"ERROR"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	OrderResponse OrderResponse `xml:"OrderResponse"`
	RouteResponse RouteResponse `xml:"RouteResponse"`
}

type OrderResponse struct {
	FilterResult string `xml:"filter_result,attr"`
	DestCode     string `xml:"destcode,attr"`
	MailNo       string `xml:"mailno,attr"`
	OriginCode   string `xml:"origincode,attr"`
	OrderId      string `xml:"orderid,attr"`
	ResStatus    string `xml:"res_status,attr"`
}

type RouteResponse struct {
	MailNo  string  `xml:"mailno,attr"`
	OrderId string  `xml:"orderid,attr"`
	Route   []Route `xml:"Route"`
}

type Route struct {
	Remark        string `xml:"remark,attr"`
	AcceptTime    string `xml:"accept_time,attr"`
	AcceptAddress string `xml:"accept_address,attr"`
	Opcode        string `xml:"opcode,attr"`
}

type RequestXml struct {
	XMLName     xml.Name    `xml:"Request"`
	Service     string      `xml:"service,attr"`
	Lang        string      `xml:"lang,attr"`
	Head        string      `xml:"Head"`
	RequestBody RequestBody `xml:"Body"`
}

type RequestBody struct {
	OrderSearch  OrderSearch  `xml:"OrderSearch"`
	OrderConfirm OrderConfirm `xml:"OrderConfirm"`
	OrderPush    OrderPush    `xml:"Order"`
	RouteRequest RouteRequest `xml:"RouteRequest"`
}

type RouteRequest struct {
	TrackingType   string `xml:"tracking_type,attr"`
	MethodType     string `xml:"method_type,attr"`
	TrackingNumber string `xml:"tracking_number,attr"`
}

type OrderPush struct {
	OrderId         string `xml:"orderid,attr"`
	ExpressType     int    `xml:"express_type,attr"`
	JCompany  string `xml:"j_company,attr"`
	JContact  string `xml:"j_contact,attr"`
	JTel      string `xml:"j_tel,attr"`
	JMobile   string `xml:"j_mobile,attr"`
	JProvince string `xml:"j_province,attr"`
	JCity     string `xml:"j_city,attr"`
	JCounty   string `xml:"j_county,attr"`
	JAddress  string `xml:"j_address,attr"`
	DCompany  string `xml:"d_company,attr"`
	DContact  string `xml:"d_contact,attr"`
	DTel      string `xml:"d_tel,attr"`
	DMobile   string `xml:"d_mobile,attr"`
	DProvince string `xml:"d_province,attr"`
	DCity     string `xml:"d_city,attr"`
	DCounty   string `xml:"d_county,attr"`
	DAddress  string `xml:"d_address,attr"`
	ParcelQuantity string       `xml:"parcel_quantity,attr"`
	PayMethod      int          `xml:"pay_method,attr"`
	Custid          string       `xml:"custid,attr"`
	customs_batchs  string       `xml:"customs_batchs,attr"`
	Cargo           string       `xml:"cargo,attr"`
	AddedService    AddedService `xml:"AddedService"`
}

type UserInfo struct {
	Company  string `xml:"d_company,attr"`
	Contact  string `xml:"d_contact,attr"`
	Tel      string `xml:"d_tel,attr"`
	Mobile   string `xml:"d_mobile,attr"`
	Province string `xml:"d_province,attr"`
	City     string `xml:"d_city,attr"`
	County   string `xml:"d_county,attr"`
	Address  string `xml:"d_address,attr"`
}

type AddedService struct {
	Name   string `xml:"name,attr"`
	Value  string `xml:"value,attr"`
	Value1 string `xml:"value1,attr"`
	Value2 string `xml:"value2,attr"`
	Value3 string `xml:"value3,attr"`
	Value4 string `xml:"value4,attr"`
	Value5 string `xml:"value5,attr"`
}

type OrderSearch struct {
	OrderId string `xml:"orderid,attr"`
}

type OrderConfirm struct {
	OrderId  string `xml:"orderid,attr"`
	DealType string `xml:"dealtype,attr"`
}

func (e Err) Error() string {
	return fmt.Sprintf("{Code: %v, Msg: %v}", e.ErrCode, e.ErrMsg)
}

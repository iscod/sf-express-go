package sfexpress

import (
	"encoding/xml"
	"fmt"
)

type ServiceName string

var (
	OrderService        ServiceName = "OrderService"        //下单接口
	OrderSearchService  ServiceName = "OrderSearchService"  //订单查询
	OrderConfirmService ServiceName = "OrderConfirmService" //订单确认和取消
	OrderFilterService  ServiceName = "OrderFilterService"  //订单筛选接口
	OrderRouteService   ServiceName = "RouteService"        //订单筛选接口
	RequestServiceLang              = "zh-CN"
	RequestServiceHead              = "SLKJ2019"
	ServiceURL                      = "http://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
	ServiceURLHttps                 = "https://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
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

//type ResponseXml struct {
//	XMLName xml.Name `xml:"Response"`
//	Head    string   `xml:"Head"`
//	Err     Err      `xml:"ERROR"`
//	Body    Body     `xml:"Body"`
//}

type Body struct {
	OrderResponse       OrderResponse       `xml:"OrderResponse"`
	RouteResponse       RouteResponse       `xml:"RouteResponse"`
	OrderFilterResponse OrderFilterResponse `xml:"OrderFilterResponse"`
}

type OrderResponse struct {
	OrderId      string `xml:"orderid,attr"`
	MailNo       string `xml:"mailno,attr"`
	OriginCode   string `xml:"origincode,attr"`
	DestCode     string `xml:"destcode,attr"`
	FilterResult string `xml:"filter_result,attr"`
	Remark       string `xml:"remark,attr"`
	ResStatus    string `xml:"res_status,attr"`
}

type RouteResponse struct {
	MailNo  string  `xml:"mailno,attr"`
	OrderId string  `xml:"orderid,attr"`
	Route   []Route `xml:"Route"`
}

type OrderFilterResponse struct {
	OrderId      string `xml:"orderid,attr"`
	FilterResult string `xml:"filter_result,attr"`
	OriginCode   string `xml:"origincode,attr"`
	Remark       string `xml:"remark,attr"`
}

type Route struct {
	Remark        string `xml:"remark,attr"`
	AcceptTime    string `xml:"accept_time,attr"`
	AcceptAddress string `xml:"accept_address,attr"`
	Opcode        string `xml:"opcode,attr"`
}

type RequestXml struct {
	XMLName     xml.Name    `xml:"Request"`
	Service     ServiceName `xml:"service,attr"`
	Lang        string      `xml:"lang,attr"`
	Head        string      `xml:"Head"`
	RequestBody RequestBody `xml:"Body"`
}

type RequestBody struct {
	OrderSearch  OrderSearch  `xml:"OrderSearch"`
	OrderConfirm OrderConfirm `xml:"OrderConfirm"`
	Order        Order        `xml:"Order"`
	RouteRequest RouteRequest `xml:"RouteRequest"`
	OrderFilter  OrderFilter  `xml:"OrderFilter"`
}

type RouteRequest struct {
	TrackingType   int    `xml:"tracking_type,attr"`   //可选，查询号类别: 1:根据顺丰运单号查询,order节点中tracking_number将被当作顺丰运单号处理 2:根据客户订单号查询,order节点中tracking_number将被当作客户订单号处理 3:逆向单,根据客户原始订单号查询,order节点中tracking_number将被当作逆向单原始订单号处理
	TrackingNumber string `xml:"tracking_number,attr"` //查询号: 如果tracking_type=1,则此值为顺丰运单号 如果tracking_type=2,则此值为客户订单号 如果tracking_type=3,则此值为逆向单原始订单号 如果有多个单号,以逗号分隔,如"123,124,125"。
	MethodType     string `xml:"method_type,attr"`     //可选，路由查询类别: 1:标准路由查询
	CheckPhoneNo   string `xml:"check_phoneNo,attr"`   //可选，校验电话号码后四位值; 按运单号查询路由时,可通过该参数传入用于校验的电话号码后4位(寄方或收方都可以),如果涉及多个运单号,对应该值也需按顺序传入多个,并以英文逗号隔开。
}

type Order struct {
	OrderId        string       `xml:"orderid,attr"` //必填
	ExpressType    int          `xml:"express_type,attr"`
	JCompany       string       `xml:"j_company,attr"`
	JContact       string       `xml:"j_contact,attr"`
	JTel           string       `xml:"j_tel,attr"`
	JMobile        string       `xml:"j_mobile,attr"`
	JProvince      string       `xml:"j_province,attr"`
	JCity          string       `xml:"j_city,attr"`
	JCounty        string       `xml:"j_county,attr"`
	JAddress       string       `xml:"j_address,attr"`
	DCompany       string       `xml:"d_company,attr"`
	DContact       string       `xml:"d_contact,attr"`
	DTel           string       `xml:"d_tel,attr"`
	DMobile        string       `xml:"d_mobile,attr"`
	DProvince      string       `xml:"d_province,attr"`
	DCity          string       `xml:"d_city,attr"`
	DCounty        string       `xml:"d_county,attr"`
	DAddress       string       `xml:"d_address,attr"`
	ParcelQuantity string       `xml:"parcel_quantity,attr"`
	PayMethod      int          `xml:"pay_method,attr"`
	CustId         string       `xml:"custid,attr"`
	CustomsBatchs  string       `xml:"customs_batchs,attr"`
	Cargo          string       `xml:"cargo,attr"`
	AddedService   AddedService `xml:"AddedService"`
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
	OrderId    string `xml:"orderid,attr"`
	SearchType int    `xml:"search_type,attr"` //可选 1,正向单查询,传入的orderid为正向定单号,2,退货单查询,传入的orderid 为退货原始订单号
}

type OrderConfirm struct {
	OrderId        string `xml:"orderid,attr"`          //客户订单号
	MailNo         string `xml:"mailno,attr"`           //顺丰母运单号如果dealtype=1,必填
	DealType       string `xml:"dealtype,attr"`         //可选，客户订单操作标识: 1:确认 2:取消
	CustomsBatchs  string `xml:"customs_batchs,attr"`   //可选，报关批次
	AgentNo        string `xml:"agent_no,attr"`         //可选，代理单号
	ConsignEmpCode string `xml:"consign_emp_code,attr"` //可选，收派员工号
}

type OrderFilter struct {
	FilterType int    `xml:"filter_type,attr"` //可选筛单类别: 1:自动筛单(系统根据地址库进行判断,并返回结果,系统无法检索到可派送的将返回不可派送) 2:可人工筛单(系统首先根据地址库判断,如果无法自动判断是否收派,系统将生成需要人工判断的任务,后续由人工处理,处理结束后,顺丰可主动推送给客户系统)
	OrderId    string `xml:"orderid,attr"`     //客户订单号,filter_type=2则必须提供
	Daddress   string `xml:"d_address,attr"`   //到件方详细地址,需要包括省市区,如:广东省深圳市福田区新洲十一街万基商务大厦。
}

func (e Err) Error() string {
	return fmt.Sprintf("{Code: %v, Msg: %v}", e.ErrCode, e.ErrMsg)
}

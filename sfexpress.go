package sfexpress

// Order 下订单接口
func (c *Config) Order(order Order) (*OrderResponse, error) {
	if order.CustId == "" && c.Custid != "" {
		order.CustId = c.Custid
	}
	result, err := c.doRequest(OrderRequestBody{Order: order})
	if err != nil {
		return nil, err
	}
	return &result.OrderResponse, nil
}

// OrderQuery 订单查询
// 因Internet环境下,网络不是绝对可靠,用户系统下订单到顺丰后,不一定可以收到顺丰系统返回的数据,此接口用于在未收到返回数据时,查询下订单(含筛选)接口客户订单当前的处理情况。
func (c *Config) OrderQuery(orderSearch OrderSearch) (*OrderResponse, error) {
	result, err := c.doRequest(OrderSearchRequestBody{OrderSearch: orderSearch})
	if err != nil {
		return nil, err
	}

	return &result.OrderResponse, nil
}

// OrderConfirm 订单确认/取消接口
// 客户在确定将货物交付给顺丰托运后,将面单上的一些重要信息,如快件重量通过此接口发送给顺丰。
// 客户在发货前取消订单
// 订单取消之后,订单号也是不能重复利用的。
func (c *Config) OrderConfirm(orderConfirm OrderConfirm) (*OrderConfirmResponse, error) {
	result, err := c.doRequest(OrderConfirmRequestBody{OrderConfirm: orderConfirm})
	if err != nil {
		return nil, err
	}
	return &result.OrderConfirmResponse, nil
}

// OrderFilterService 客户系统通过此接口向顺丰系统发送主动的筛单请求,用于判断客户的收、派地址是否属于顺丰的收派范围。
func (c *Config) OrderFilterService(orderFilter OrderFilter) (*OrderFilterResponse, error) {
	result, err := c.doRequest(OrderFilterRequestBody{OrderFilter: orderFilter})
	if err != nil {
		return nil, err
	}
	return &result.OrderFilterResponse, nil
}

//OrderRouteService 路由查询接口
//客户可通过此接口查询顺丰运单路由,系统将返回当前时间点已发生的路由信息。
//路由查询接口支持三种查询方式:
//根据通过丰桥接口下单的订单号查询,系统校验信息匹配将返回对应运单路由信息。
//根据运单号+月结卡号(需与传入的顾客编码绑定)查询,系统校验信息归属关系正确将返回对应运单路由信息。
//根据运单号+运单对应的收寄人任一方电话号码后4位(参数check_phoneNo中传入)查询,系统校验信息匹配将返回对应运单路由信息。
func (c *Config) OrderRouteService(routeRequest RouteRequest) ([]RouteResponse, error) {
	result, err := c.doRequest(RouteRequestBody{RouteRequest: routeRequest})
	if err != nil {
		return nil, err
	}
	return result.RouteResponse, nil
}

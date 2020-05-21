# Go Sf-Express

Welcome to use sf-express-go

## Installation

```sh
go get -u -v github.com/IsCod/sf-express-go
```

## Quickstart

```go
package main

import (
	"fmt"
    "math/rand"
	sfexpress "sf-express-go"
	"time"
)

func main() {
	c := sfexpress.Config{CheckWord: "FBIqMkZjzxbsZgo7jTpeq7PD8CVzLT4Q", ClientCode: "NTL", Custid: "7553032834"}

	orderId := "SF-" + time.Now().Format("20060102150405") + string(rand.Intn(100))

	//order push
	o, err := c.OrderPush(orderId, sfexpress.UserInfo{
		Company:  "顺丰镖局",
		Contact:  "发件人",
		Tel:      "15012345678",
		Mobile:   "",
		Province: "广东省",
		City:     "深圳市",
		County:   "南山区",
		Address:  "福田区新洲十一街万基商务大厦26楼",
	}, sfexpress.UserInfo{
		Company:  "顺丰镖局",
		Contact:  "收件人",
		Tel:      "15012345678",
		Mobile:   "",
		Province: "广东省",
		City:     "深圳市",
		County:   "南山区",
		Address:  "福田区新洲十一街万基商务大厦26楼",
	}, 1, 1, sfexpress.AddedService{
		Name:  "COD",
		Value: "1.01",
	}, "iphone x")
	if err != nil {
		fmt.Printf("❌ Push: %s\n", err.Error())
	} else {
		fmt.Printf("✅ Push: orderId: %s, MailNo: %s\n", o.OrderId, o.MailNo)
	}

	//order query
	order, err := c.OrderQuery("QIAO-20171231001")
	if err != nil {
		fmt.Printf("❌ query: %s\n", err)
	} else {
		fmt.Printf("✅ query: orderId: %s, MailNo: %s\n", order.OrderId, order.MailNo)
	}

	//order confirm&cancel
	oc, err := c.OrderConfirm("XJFS_071100251", "2")
	if err != nil {
		fmt.Printf("❌ Confirm: %s\n", err)
	} else {
		fmt.Printf("✅ Confirm:  orderId: %s, MailNo: %s\n", oc.OrderId, oc.ResStatus)
	}

	//route Query By OrderNo
	oro, err := c.OrderRouteQueryByOrderNo("XJFS_071100251")
	if err != nil {
		fmt.Printf("❌ RouteQueryByOrderNo: %s", err)
	} else {
		fmt.Printf("✅ RouteQueryByOrderNo:  orderId: %s, MailNo: %s, %v", oro.OrderId, oro.MailNo, oro.Route)
	}
}

```

## Development
TODO:

## Contributing


## License

The sf-express-go is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

## Code of Conduct

Everyone interacting in the SfExpress project’s codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](https://github.com/[USERNAME]/sf_express/blob/master/CODE_OF_CONDUCT.md).

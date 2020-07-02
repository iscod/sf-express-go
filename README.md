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
	o, err := c.Order(sfexpress.Order{
		OrderId:     orderId,
		ExpressType: 1,
		PayMethod:   1,
		JCompany:    "顺丰镖局",
		JContact:    "虾哥",
		JTel:        "15012345678",
		JMobile:     "",
		JProvince:   "广东省",
		JCity:       "深圳市",
		JCounty:     "南山区",
		JAddress:    "福田区新洲十一街万基商务大厦26楼",
		DCompany:    "顺丰镖局",
		DContact:    "虾哥",
		DTel:        "15012345678",
		DMobile:     "",
		DProvince:   "广东省",
		DCity:       "深圳市",
		DCounty:     "南山区",
		DAddress:    "福田区新洲十一街万基商务大厦26楼",
		AddedService: sfexpress.AddedService{
			Name:  "COD",
			Value: "1.01",
		},
		Cargo: "iphone x",
	})
	if err != nil {
		fmt.Printf("❌ Push: %s\n", err.Error())
	} else {
		fmt.Printf("✅ Push: orderId: %s, MailNo: %s\n", o.OrderId, o.MailNo)
	}

	//order query
	order, err := c.OrderQuery(sfexpress.OrderSearch{OrderId: o.OrderId, SearchType: 1})
	if err != nil {
		fmt.Printf("❌ query: %s\n", err)
	} else {
		fmt.Printf("✅ query: orderId: %s, MailNo: %s\n", order.OrderId, order.MailNo)
	}

	//order confirm&cancel
	oc, err := c.OrderConfirm(sfexpress.OrderConfirm{OrderId: o.OrderId, DealType: "2"})
	if err != nil {
		fmt.Printf("❌ Confirm: %s\n", err)
	} else {
		fmt.Printf("✅ Confirm:  orderId: %s, ResStatus: %d\n", oc.OrderId, oc.ResStatus)
	}

	//route Query By OrderNo
	oro, err := c.OrderRouteService(sfexpress.RouteRequest{TrackingType: 1, TrackingNumber: "XJFS_071100251"})
	if err != nil {
		fmt.Printf("❌ RouteQueryByOrderNo: %s", err)
	} else {
		for _, i := range oro {
			fmt.Printf("✅ RouteQueryByOrderNo:  orderId: %s, MailNo: %s, %v", i.OrderId, i.MailNo, i.Route)
		}
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

package main

import (
	"fmt"
	"time"
)

type shop struct{
	name string
	url string
}

type order struct{
	id int
	amount float32
	customerName string
	status string
	shop
	createdAt time.Time
}

func newOrder (id int, amount float32, customerName string, status string) *order {
	myOrder := order {
		id : id,
		amount: amount,
		customerName: customerName,
		status: status,
	}	

	return &myOrder
}

func (o *order)changeStatus (status string){
	o.status = status
}

func (o order)getAmount() float32{
	return o.amount
}

func main(){
	shop1 := shop{
		name: "New Shop",
		url: "new.shop.com",
	}

	order1 := order {
		id : 1,
		amount: 10.6,
		customerName: "Farhan",
		shop: shop1,
	}	

	order2 := order {
		id : 2,
		amount: 32.6,
		customerName: "Ahmed",
		shop: shop{
			name: "shop2",
			url: "new.shop2.com",
		},
		createdAt: time.Now(),
	}

	order3 := newOrder(3, 67.89, "Nahid", "")

	order1.createdAt = time.Now()
	order1.changeStatus("PAID")
	fmt.Println(order2.getAmount())
	fmt.Println(order1.customerName)
	fmt.Println(order1)
	fmt.Println(order2)
	fmt.Println(order3)

	language := struct {
		name  string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
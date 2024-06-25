package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 1)
	sellOrderA := NewOrder(false, 2)
	sellOrderB := NewOrder(false, 3)
	sellOrderC := NewOrder(false, 4)

	l.AddOrder(buyOrderA)
	l.AddOrder(sellOrderA)
	l.AddOrder(sellOrderB)
	l.AddOrder(sellOrderC)

	l.DeleteOrder(sellOrderA)

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 15)
	buyOrderB := NewOrder(true, 2_000)

	ob.PlaceOrder(18_000, buyOrderA)
	ob.PlaceOrder(18_000, buyOrderB)

	fmt.Printf("%+v", ob)
}

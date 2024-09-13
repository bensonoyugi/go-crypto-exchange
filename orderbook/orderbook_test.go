package orderbook

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

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

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrderA := NewOrder(false, 10)
	sellOrderB := NewOrder(false, 12)
	buyOrderA := NewOrder(true, 100)

	ob.PlaceLimitOrder(10_000, sellOrderA)
	ob.PlaceLimitOrder(10_000, sellOrderB)
	ob.PlaceLimitOrder(10_000, buyOrderA)

	assert(t, len(ob.AskLimits[10_000].Orders), 2)
	assert(t, len(ob.bids), 1)
}

func TestPlaceMarketOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder := NewOrder(false, 20)
	ob.PlaceLimitOrder(10_000, sellOrder)

	buyOrder := NewOrder(true, 10)
	matches := ob.PlaceMarketOrder(buyOrder)

	assert(t, len(matches), 1)
	assert(t, len(ob.asks), 1)
	assert(t, ob.AskTotalVolume(), 10.0)
	assert(t, matches[0].Ask, sellOrder)
	assert(t, matches[0].Bid, buyOrder)
	assert(t, matches[0].SizeFilled, 10.0)
	assert(t, matches[0].Price, 10_000.0)
	assert(t, buyOrder.IsFilled(), true)
}

func TestPlaceMarketOrderMultiFill(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)
	buyOrderD := NewOrder(true, 1)

	ob.PlaceLimitOrder(10_000, buyOrderA)
	ob.PlaceLimitOrder(5_000, buyOrderB)
	ob.PlaceLimitOrder(9_000, buyOrderC)
	ob.PlaceLimitOrder(6_000, buyOrderD)

	assert(t, ob.BidTotalVolume(), 24.0)

	sellOrder := NewOrder(false, 20)
	matches := ob.PlaceMarketOrder(sellOrder)

	assert(t, ob.BidTotalVolume(), 4.0)
	assert(t, len(matches), 4)
	assert(t, len(ob.bids), 1)

	fmt.Printf("%+v", matches)
}

func TestCancelOrder(t *testing.T) {
	ob := NewOrderbook()

	buyOrder := NewOrder(true, 4)

	ob.PlaceLimitOrder(10_000, buyOrder)

	assert(t, ob.BidTotalVolume(), 4.0)

	ob.CancelOrder(buyOrder)

	assert(t, ob.BidTotalVolume(), 0.0)
}

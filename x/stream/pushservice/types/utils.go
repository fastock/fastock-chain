package types

import "github.com/fastock/fastock-chain/x/backend"

func find(orders []backend.Order, o backend.Order) (i int, found bool) {
	for i, ord := range orders {
		if ord.OrderID == o.OrderID {
			return i, true
		}
	}
	return -1, false
}

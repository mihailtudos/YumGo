package module

import "github.com/mihailtudos/yumgo/backend/orders/api/module/client"

type Orders struct{}

func (o Orders) PingOrders(request client.PingOrdersRequest) error {
	return nil
}

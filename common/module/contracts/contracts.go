package contracts

import (
	"errors"

	ordersModule "github.com/mihailtudos/yumgo/backend/orders/api/module/client"
)

type Contracts struct {
	ordersModule.Orders
}

func (c *Contracts) Verify() error {
	var err error

	if c.Orders == nil {
		err = errors.Join(err, errors.New("orders module contract is empty"))
	}
	return err
}

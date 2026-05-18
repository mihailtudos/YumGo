package module

import (
	"context"

	"github.com/mihailtudos/yumgo/common"
	"github.com/mihailtudos/yumgo/common/module/contracts"
)

type Name string

type Module interface {
	Name() Name
	Init(ctx context.Context) error
	RegisterHTTP(ctx context.Context, e common.EchoRouter) error
	RegisterContracts(ctx context.Context, contracts *contracts.Contracts) error
}

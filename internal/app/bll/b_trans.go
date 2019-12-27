package bll

import (
	"context"
	"github.com/authhwang/gin-admin/internal/app/bll/impl/internal"

	"github.com/google/wire"
)

// ITrans 事务管理接口
type ITrans interface {
	// 执行事务
	Exec(ctx context.Context, fn func(context.Context) error) error
}

var BllTransSet = wire.NewSet(
	internal.NewTrans,
	wire.Bind(new(ITrans), new(*internal.Trans)),
)

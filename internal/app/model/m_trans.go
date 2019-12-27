package model

import (
	"context"

	"github.com/authhwang/gin-admin/internal/app/model/impl/gorm/internal/model"

	"github.com/google/wire"
)

// ITrans 事务管理接口
type ITrans interface {
	// 开始事务
	Begin(ctx context.Context) (interface{}, error)
	// 提交事务
	Commit(ctx context.Context, trans interface{}) error
	// 回滚事务
	Rollback(ctx context.Context, trans interface{}) error
}

var ModelTransSet = wire.NewSet(
	model.NewTrans,
	wire.Bind(new(ITrans), new(*model.Trans)),
)

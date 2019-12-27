package model

import (
	"context"

	"github.com/authhwang/gin-admin/internal/app/schema"
	"github.com/authhwang/gin-admin/internal/app/model/impl/gorm/internal/model"

	"github.com/google/wire"
)

// IMenu 菜单管理存储接口
type IMenu interface {
	// 查询数据
	Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.MenuQueryOptions) (*schema.Menu, error)
	// 创建数据
	Create(ctx context.Context, item schema.Menu) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Menu) error
	// 更新父级路径
	UpdateParentPath(ctx context.Context, recordID, parentPath string) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}

var ModelMenuSet = wire.NewSet(
	model.NewMenu,
	wire.Bind(new(IMenu), new(*model.Menu)),
)

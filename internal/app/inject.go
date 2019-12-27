//+build wireinject
package app

import (
	"context"
	
	"github.com/authhwang/gin-admin/internal/app/config"
	
	"github.com/authhwang/gin-admin/internal/app/routers/api/ctl"
	"github.com/authhwang/gin-admin/internal/app/bll"
	"github.com/authhwang/gin-admin/internal/app/model"

	"github.com/jinzhu/gorm"
	"github.com/casbin/casbin"
	"github.com/google/wire"
)

var c *casbin.Enforcer

func NewEnforcer() *casbin.Enforcer {
	cfg := config.Global()
	if c == nil {
		c = casbin.NewEnforcer(cfg.CasbinModelConf, false)
	}
	return c
}

func InitDataStore(ctx context.Context, db *gorm.DB) error {
	wire.Build(
		NewEnforcer,
		model.ModelTransSet,
		model.ModelUserSet,
		model.ModelMenuSet,
		model.ModelRoleSet,
		bll.BllUserSet,
		bll.BllMenuSet,
		bll.BllRoleSet,
		bll.BllTransSet,
		InitData,
	)

	return nil
}

func InitWebFunc(ctx context.Context, db *gorm.DB) (func(), error) {
	wire.Build(
		model.ModelTransSet,
		model.ModelUserSet,
		model.ModelDemoSet,
		model.ModelMenuSet,
		model.ModelRoleSet,
		bll.BllDemoSet,
		bll.BllUserSet,
		bll.BllMenuSet,
		bll.BllRoleSet,
		bll.BllLoginSet,
		ctl.CtlSet,
		NewEnforcer,
		InitAuth,
		InitHTTPServer,
	)

	return nil, nil
}


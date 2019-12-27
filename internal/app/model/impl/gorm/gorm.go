package gorm

import (
	"time"

	"github.com/authhwang/gin-admin/internal/app/config"
	//"github.com/authhwang/gin-admin/internal/app/model"
	"github.com/authhwang/gin-admin/internal/app/model/impl/gorm/internal/entity"
	//imodel "github.com/authhwang/gin-admin/internal/app/model/impl/gorm/internal/model"
	"github.com/jinzhu/gorm"
	//"github.com/google/wire"
	//"go.uber.org/dig"

	// gorm存储注入
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Config 配置参数
type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

// NewDB 创建DB实例
func NewDB(c *Config) (*gorm.DB, error) {
	db, err := gorm.Open(c.DBType, c.DSN)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)
	return db, nil
}

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	entity.SetTablePrefix(prefix)
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.Global().Gorm.DBType; dbType == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(entity.Demo),
		new(entity.User),
		new(entity.UserRole),
		new(entity.Role),
		new(entity.RoleMenu),
		new(entity.Menu),
		new(entity.MenuAction),
		new(entity.MenuResource),
	).Error
}

// Inject 注入gorm实现
// 使用方式：
//   container := dig.New()
//   Inject(container)
//   container.Invoke(func(foo IDemo) {
//   })
// func Inject(container *dig.Container) error {
// 	container.Provide(imodel.NewTrans, dig.As(new(model.ITrans)))
// 	container.Provide(imodel.NewDemo, dig.As(new(model.IDemo)))
// 	container.Provide(imodel.NewMenu, dig.As(new(model.IMenu)))
// 	container.Provide(imodel.NewRole, dig.As(new(model.IRole)))
// 	container.Provide(imodel.NewUser, dig.As(new(model.IUser)))
// 	return nil
// }

// var GormSet = wire.NewSet(
// 	imodel.NewTrans,
// 	imodel.NewDemo,
// 	imodel.NewMenu,
// 	imodel.NewRole,
// 	imodel.NewUser,
// 	wire.Bind(new(model.ITrans), new(*imodel.Trans)),
// 	wire.Bind(new(model.IDemo), new(*imodel.Demo)),
// 	wire.Bind(new(model.IMenu), new(*imodel.Menu)),
// 	wire.Bind(new(model.IRole), new(*imodel.Role)),
// 	wire.Bind(new(model.IUser), new(*imodel.User)),
// ) 
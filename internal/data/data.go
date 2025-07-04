package data

import (
	"realworld_demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers. 需要加载的要加入到这里
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewGreeterRepo,
	NewUserRepo,
	NewProfileRepo,
	NewArticleRepo,
	NewCommentRepo,
)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	log.Info("Connecting to database...")
	log.Infof("DSN: %s", c.Database.Dsn)
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	InitDB(db)
	return db
}

func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(
		&User{},
		&Article{},
		&Comment{},
		&ArticleFavorite{},
		&Following{},
	); err != nil {
		panic(err)
	}
}

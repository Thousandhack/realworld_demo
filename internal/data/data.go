package data

import (
	"fmt"
	"realworld_demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers. 依赖注入的集合
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

	// 尝试连接数据库
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Errorf("数据库连接失败: %v", err)
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// 测试数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("获取数据库连接失败: %v", err)
		panic(fmt.Sprintf("failed to get database connection: %v", err))
	}

	if err := sqlDB.Ping(); err != nil {
		log.Errorf("数据库Ping失败: %v", err)
		panic(fmt.Sprintf("database ping failed: %v", err))
	}

	log.Info("数据库连接成功")

	// 初始化数据库表
	InitDB(db)
	return db
}

func InitDB(db *gorm.DB) {
	log.Info("开始初始化数据库表...")

	// 检查数据库表是否存在
	tables := []string{"users", "articles", "comments", "article_favorites", "follow_users"}
	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			log.Infof("表 %s 已存在", table)
		} else {
			log.Infof("表 %s 不存在，将创建", table)
		}
	}

	// 自动迁移表结构
	if err := db.AutoMigrate(
		&User{},
		&Article{},
		&Comment{},
		&ArticleFavorite{},
		&FollowUser{},
	); err != nil {
		log.Errorf("数据库表自动迁移失败: %v", err)
		panic(fmt.Sprintf("database migration failed: %v", err))
	}

	log.Info("数据库表初始化完成")
}

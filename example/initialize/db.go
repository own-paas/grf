package initialize

import (
	"github.com/sestack/grf/example/global"
	"github.com/sestack/grf/example/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       global.CONFIG.DB.DSN, // DSN data source name
		DefaultStringSize:         128,                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return err
	}
	return nil
}

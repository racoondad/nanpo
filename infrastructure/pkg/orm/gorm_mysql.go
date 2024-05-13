package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ensureMysql(conf IGormConfig) error {
	mysqlConfig := mysql.Config{
		DSN:                       conf.MainDsn(), // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return err
	} else {
		return db.Exec(conf.CreateProvider()).Error
	}
}

func gormMysqlByConfig(conf IGormConfig) *gorm.DB {
	if conf.DbnameParam() == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       conf.Dsn(), // DSN data source name
		DefaultStringSize:         191,        // string 类型字段的默认长度
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormImpl.Config(conf)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConnsParam())
		sqlDB.SetMaxOpenConns(conf.MaxOpenConnsParam())
		return db
	}
}

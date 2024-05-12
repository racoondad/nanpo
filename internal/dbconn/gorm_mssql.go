package dbconn

import (
	"github.com/racoondad/nanpo/internal/configs"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ensureMssql(conf configs.SpecializedDB) error {
	mssqlConfig := sqlserver.Config{
		DSN:               conf.MainDsn(), // DSN data source name
		DefaultStringSize: 191,            // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig)); err != nil {
		return err
	} else {
		return db.Exec(conf.CreateProvider()).Error
	}
}

func gormMssqlByConfig(conf configs.SpecializedDB) *gorm.DB {
	if conf.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               conf.Dsn(), // DSN data source name
		DefaultStringSize: 191,        // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), gormImpl.Config(conf)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		return db
	}
}

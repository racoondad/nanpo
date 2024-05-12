package dbconn

// import (
// 	"saasbackend/configs"

// 	"github.com/dzwvip/oracle"
// 	"gorm.io/gorm"
// )

// func ensureOracle(conf configs.SpecializedDB) error {
// 	oracleConfig := oracle.Config{
// 		DSN:               conf.MainDsn(), // DSN data source name
// 		DefaultStringSize: 191,            // string 类型字段的默认长度

// 	}
// 	if db, err := gorm.Open(oracle.New(oracleConfig)); err != nil {
// 		return err
// 	} else {
// 		return db.Exec(conf.CreateProvider()).Error
// 	}
// }

// func gormOracleByConfig(conf configs.SpecializedDB) *gorm.DB {
// 	if conf.Dbname == "" {
// 		return nil
// 	}
// 	oracleConfig := oracle.Config{
// 		DSN:               conf.Dsn(), // DSN data source name
// 		DefaultStringSize: 191,        // string 类型字段的默认长度

// 	}
// 	if db, err := gorm.Open(oracle.New(oracleConfig), gormImpl.Config(conf)); err != nil {
// 		return nil
// 	} else {
// 		sqlDB, _ := db.DB()
// 		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
// 		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
// 		return db
// 	}
// }

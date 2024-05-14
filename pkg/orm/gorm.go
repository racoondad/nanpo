package orm

import (
	"fmt"
	"strings"

	"github.com/racoondad/nanpo/configs"
	"gorm.io/gorm"
)

func GormConfig(conf configs.SpecializedDB) (db *gorm.DB, err error) {
	switch strings.ToLower(conf.DbType) {
	case "mssql", "sqlserver", "sql server":
		err = ensureMssql(conf)
		if err != nil {
			return nil, err
		}
		db = gormMssqlByConfig(conf)
	case "pgsql", "postgres", "postgresql":
		ensurePgsql(conf)
		db = gormPgsqlByConfig(conf)
	// case "oracle", "orcl":
	// 	db = gormPgsqlByConfig(conf)
	case "mysql", "mariadb", "maria db":
		err = ensureMysql(conf)
		if err != nil {
			return nil, err
		}
		db = gormMysqlByConfig(conf)
	}

	if db == nil {
		err = fmt.Errorf("faild connect database by config %v", conf)
	}

	return
}

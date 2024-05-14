/*
 * @Author       : lptecodad lptecodad@sina.com
 * @Date         : 2022-08-28 00:09:04
 * @LastEditors: lptecodad lptecodad@sina.com
 * Copyright (c) 2022 by lptecodad lptecodad@sina.com, All Rights Reserved.
 */
package configs

import (
	"fmt"
)

type SpecializedDB struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`               // 数据库类型
	Host         string `mapstructure:"host" yaml:"host"`                                    // 服务器地址:端口
	Port         uint   `mapstructure:"port" yaml:"port"`                                    // 端口
	Config       string `mapstructure:"config" yaml:"config"`                                // 高级配置
	Dbname       string `mapstructure:"dbname" yaml:"dbname"`                                // 数据库名
	Username     string `mapstructure:"username" yaml:"username"`                            // 数据库用户名
	Password     string `mapstructure:"password" yaml:"password"`                            // 数据库密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                  // 全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`            // 是否开启全局禁用复数，true表示开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"` // 数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`                // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"`                // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" yaml:"log-mode"`                            // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" yaml:"log-zap"`                              // 是否通过zap写入日志文件
}

func (i *SpecializedDB) MainDsn() string {
	switch i.DbType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", i.Username, i.Password, i.Host, i.Port, "mysql", i.Config)
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s%s", i.Username, i.Password, i.Host, i.Port, "master", i.Config)
	case "oracle":
		return fmt.Sprintf("oracle://%s:%s@%s:%d/%s?%s", i.Username, i.Password, i.Host, i.Port, "orcl", i.Config)
	default:
		return fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s %s", i.Host, i.Username, i.Password, i.Port, "postgres", i.Config)
	}
}

func (i *SpecializedDB) Dsn() string {
	switch i.DbType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", i.Username, i.Password, i.Host, i.Port, i.Dbname, i.Config)
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s%s", i.Username, i.Password, i.Host, i.Port, i.Dbname, i.Config)
	case "oracle":
		return fmt.Sprintf("oracle://%s:%s@%s:%d/%s?%s", i.Username, i.Password, i.Host, i.Port, i.Dbname, i.Config)
	default:
		return fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s %s", i.Host, i.Username, i.Password, i.Port, i.Dbname, i.Config)
	}
}

func (i SpecializedDB) CreateProvider() string {
	switch i.DbType {
	case "mssql":
		return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s COLLATE Chinese_PRC_CI_AS;", i.Dbname)
	case "mysql":
		return fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", i.Dbname)
	case "pgsql":
		return fmt.Sprintf("CREATE DATABASE %s;", i.Dbname)
	case "oracle":
		return fmt.Sprintf("CREATE DATABASE %s;", i.Dbname)
	default:
		return ""
	}
}

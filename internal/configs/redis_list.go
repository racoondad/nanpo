/*
 * @Author       : lptecodad lptecodad@sina.com
 * @Date         : 2022-09-18 00:33:12
 * @LastEditors  : lptecodad lptecodad@sina.com
 * Copyright (c) 2022 by lptecodad lptecodad@sina.com, All Rights Reserved.
 */
package configs

type SpecializedCache struct {
	Zoon     int    `mapstructure:"zoon" yaml:"zoon"`         // redis的哪个数据库
	Host     string `mapstructure:"host" yaml:"host"`         // 服务器地址:端口
	Password string `mapstructure:"password" yaml:"password"` // 密码
}

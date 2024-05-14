/*
 * @Author       : lptecodad lptecodad@sina.com
 * @Date         : 2022-10-18 09:45:56
 * @LastEditors  : lptecodad lptecodad@sina.com
 * Copyright (c) 2022 by lptecodad lptecodad@sina.com, All Rights Reserved.
 */
package configs

import "time"

type Server struct {
	UseMultipoint     bool          `mapstructure:"use-multipoint" yaml:"use-multipoint"`           // 多点登录拦截
	UseTenant         bool          `mapstructure:"use-tenant" yaml:"use-tenant"`                   // 使用多租户
	Env               string        `mapstructure:"env" yaml:"env"`                                 // 环境值
	Addr              string        `mapstructure:"addr" yaml:"addr"`                               // 端口值
	ReadTimeout       time.Duration `mapstructure:"read-timeout" yaml:"read-timeout"`               // 读取超时
	ReadHeaderTimeout time.Duration `mapstructure:"read-header-timeout" yaml:"read-header-timeout"` // 读取超时
	WriteTimeout      time.Duration `mapstructure:"write-timeout" yaml:"write-timeout"`             // 写入超时
	IdleTimeout       time.Duration `mapstructure:"idle-timeout" yaml:"idle-timeout"`               // 提交超时
	MaxHeaderBytes    int           `mapstructure:"max-header-bytes" yaml:"max-header-bytes"`       // 请求头最大值

	LimitIP  int `mapstructure:"limit-ip" yaml:"limit-ip"`   // 限制IP次数
	LimitJWT int `mapstructure:"limit-jwt" yaml:"limit-jwt"` // 限制接口调用次数
}

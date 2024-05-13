/*
 * @Author       : lptecodad lptecodad@sina.com
 * @Date         : 2022-08-28 00:09:04
 * @LastEditors: lptecodad lptecodad@sina.com
 * Copyright (c) 2022 by lptecodad lptecodad@sina.com, All Rights Reserved.
 */
package configs

import (
	"reflect"

	"github.com/spf13/viper"
)

type Server struct {
	JWT        JWT     `mapstructure:"jwt" yaml:"jwt"`
	Zap        Zap     `mapstructure:"zap" yaml:"zap"`
	System     System  `mapstructure:"system" yaml:"system"`
	Captcha    Captcha `mapstructure:"captcha" yaml:"captcha"`
	MainTenant Tenant  `mapstructure:"main-tenant" yaml:"main-tenant"`
}

func WriteConfig(conf Server, vp *viper.Viper) (err error) {
	cs := structToMapstructure(conf)
	for k, v := range cs {
		vp.Set(k, v)
	}
	err = vp.WriteConfig()
	if err != nil {
		return
	}

	return vp.ReadInConfig()
}

func structToMapstructure(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

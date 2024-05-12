package configs

// 微信小程序
type WxMini struct {
	AppId   string `mapstructure:"app-id" yaml:"app-id"`   //
	Secret  string `mapstructure:"secret" yaml:"secret"`   //
	Version string `mapstructure:"version" yaml:"version"` // 小程序版本号
}

type WxPub struct{}

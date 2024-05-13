package configs

type Captcha struct {
	KeyLong    int `mapstructure:"key-long" yaml:"key-long"`       // 验证码长度
	CapchaType int `mapstructure:"capcha-type" yaml:"capcha-type"` // 验证码类型
	ImgWidth   int `mapstructure:"img-width" yaml:"img-width"`     // 验证码宽度
	ImgHeight  int `mapstructure:"img-height" yaml:"img-height"`   // 验证码高度
}

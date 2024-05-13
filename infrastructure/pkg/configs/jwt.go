package configs

type JWT struct {
	SigningKey string `mapstructure:"signing-key" yaml:"signing-key"` // jwt签名
	Expire     int    `mapstructure:"expire" yaml:"expire"`           // 有效分钟
	Buffer     int64  `mapstructure:"buffer" yaml:"buffer"`           // 刷新分钟
	Issuer     string `mapstructure:"issuer" yaml:"issuer"`           // 签发者
}

func (j JWT) BufferAt() int64 {
	return j.Buffer
}

func (j JWT) GetIssuer() string {
	return j.Issuer
}

func (j JWT) GetExpire() int {
	return j.Expire
}

func (j JWT) GetSigningKey() []byte {
	return []byte(j.SigningKey)
}

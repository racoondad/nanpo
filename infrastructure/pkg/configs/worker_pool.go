package configs

type WorkerPool struct {
	MaxIdle int `mapstructure:"max_idle" yaml:"max_idle"`
	MaxOpen int `mapstructure:"max_open" yaml:"max_open"`
}

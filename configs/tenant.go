package configs

type Tenant struct {
	TenantId   string           `mapstructure:"tenand-id" yaml:"tenand-id"`
	TenantName string           `mapstructure:"tenant-name" yaml:"tenant-name"`
	Enable     bool             `mapstructure:"enable" yaml:"enable"`
	Database   SpecializedDB    `mapstructure:"database" yaml:"database"`
	Redis      SpecializedCache `mapstructure:"redis" yaml:"redis"`
}

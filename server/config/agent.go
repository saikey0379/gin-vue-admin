package config

type Agent struct {
	Port      int    `mapstructure:"port" json:"port" yaml:"port"`                   // Agent端口
	CaCrt     string `mapstructure:"ca-crt" json:"ca-crt" yaml:"ca-crt"`             // certFile路径
	ClientCrt string `mapstructure:"client-crt" json:"client-crt" yaml:"client-crt"` // certFile路径
	ClientKey string `mapstructure:"client-key" json:"client-crt" yaml:"client-crt"` // certFile路径
}

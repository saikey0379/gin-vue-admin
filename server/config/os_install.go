package config

type OsInstall struct {
	DirPxe string `mapstructure:"dir-pxe" json:"dir-pxe" yaml:"dir-pxe"` // Agent端口
}

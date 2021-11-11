package config

type GlobalConfig struct {
	DB   DBConfig   `json:"db"`
	Http HttpConfig `json:"http"`
}

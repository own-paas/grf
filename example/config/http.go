package config

type HttpConfig struct {
	SSL     bool   `json:"ssl"`
	Address string `json:"address"`
}

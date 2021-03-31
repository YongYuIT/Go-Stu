package config

type DBConfig struct {
	Conn string `config:"conn"`
	Type string `config:"type"`
}

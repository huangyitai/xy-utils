package srvx

// Config TODO
type Config struct {
	Service     string `mapstructure:"service"`
	Instance    string `mapstructure:"instance"`
	Platform    string `mapstructure:"platform"`
	Project     string `mapstructure:"project"`
	Environment string `mapstructure:"environment"`
	Binary      string `mapstructure:"binary"`
}

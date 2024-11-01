package configs

import "github.com/spf13/viper"

var config *Config

type Option struct {
	ConfigFolder string
	ConfigFile   string
	ConfigType   string
}

func Init(opt Option) error {
	viper.AddConfigPath(opt.ConfigFolder)
	viper.SetConfigName(opt.ConfigFile)
	viper.SetConfigType(opt.ConfigType)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}

	return config
}

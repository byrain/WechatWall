package common

import "github.com/BurntSushi/toml"

type Configs struct {
	Listen string
}

var Config *Configs

func InitConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	Config = config
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}

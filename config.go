package config

import (
    toml "github.com/pelletier/go-toml"
    "github.com/yaotian/logs"
)

var Config *toml.TomlTree

func init() {
    config_toml, err := toml.LoadFile("conf/config.toml")
    if err != nil {
        logs.Logger.Error("Error ", err.Error())
    }
    Config = config_toml
}

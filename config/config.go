package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	conf *Config
)

// Config 全ての設定を格納
type (
	Config struct {
		Discord Discord
		Cmd     Cmd
	}
	// Discord Discordの設定
	Discord struct {
		Token    string `toml:"Token"`
		ClientID string `toml:"ClientID"`
		Prefix   string `toml:"Prefix"`
	}
	// Cmd コマンドの設定
	Cmd struct {
		Dir string `toml:"Dir"`
	}
)

func init() {
	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		f = "config/config.toml"
	}

	_, err := toml.DecodeFile(f, &conf)
	if err != nil {
		fmt.Println("設定を読み込めませんでした: ", err)
	}
}

// GetConf 全ての設定を返す
func GetConf() *Config {
	return conf
}

// GetPrefix prefixを返す
func GetPrefix() string {
	return conf.Discord.Prefix
}

// GetCmdDir コマンドディレクトリを返す
func GetCmdDir() string {
	return conf.Cmd.Dir
}

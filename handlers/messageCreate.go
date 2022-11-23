// Package handlers :イベントが起きた際に呼び出される関数をまとめたもの
package handlers

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/VegetableMeat/discordGoDemo/cmds"
	"github.com/VegetableMeat/discordGoDemo/common"
	"github.com/VegetableMeat/discordGoDemo/config"
	"github.com/bwmarrin/discordgo"
)

// Files :コマンドファイル名一覧が格納されている
// Funcs :Package cmdsの関数が格納されている
var (
	Files []string
	Funcs = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, array []string){
		"test": cmds.Test}
)

func init() {
	loadCommands()
}

// MessageCreate :BOTに対し指定したコマンドを実行する
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// BOTが自分の発言に反応しないように
	if m.Author.ID == s.State.User.ID {
		return
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 入力されたコマンドが存在していればそれを実行する
	content, arg := common.SplitContent(m.Content)
	for _, path := range Files {
		if content == config.GetPrefix()+path {
			if ch.Type != discordgo.ChannelTypeGuildText {
				s.ChannelMessageSend(m.ChannelID, "ここではコマンドを使用出来ません")
				return
			}
			Funcs[path](s, m, arg)
		}
	}
}

// loadCommands :コマンドファイルの一覧を取得してファイル名を配列へ格納
func loadCommands() {
	files, err := ioutil.ReadDir(config.GetCmdDir())
	if err != nil {
		fmt.Println("コマンドを取得できませんでした: ", err)
		return
	}

	for _, file := range files {
		fName := file.Name()
		if !file.IsDir() && path.Ext(fName) == ".go" {
			Files = append(Files, filepath.Base(fName[:len(fName)-len(filepath.Ext(fName))]))
		}
	}
}

// +build darwin,amd64 windows linux,!android
// +build go1.1

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"タイトル/utils"

	"github.com/jinzhu/gorm"
)

type ConfigList struct {
	Port    string `json:"port"`
	LogFile string `json:"log_file"`
	View    string `json:"view"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// Config Configの定義
var Config ConfigList

var Db *gorm.DB

// ポート手動変更させるためここで定義
var FlagPort string

func init() {
	// 設定ファイルconfigの読み込み
	err := LoadConfig()
	if err != nil {
		log.Printf("ファイルの読み込みに失敗しました: %v", err)
		os.Exit(1)
	}

	// ログファイルの設定
	utils.LoggingSettings(Config.LogFile)

	// コマンドの実行
	err = utils.Command()
	if err != nil {
		log.Println(err)
	}

}

// LoadConfig Configの設定
func LoadConfig() error {
	f, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, &Config)
	if err != nil {
		return err
	}

	format := "Port: %s\nLogFile: %s\nView: %s\nURL: %s\nVersion: %s\n"
	_, err = fmt.Printf(format, Config.Port, Config.LogFile, Config.View, Config.URL, Config.Version)
	if err != nil {
		return err
	}

	return nil
}

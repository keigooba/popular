package main

import (
	"flag"
	"fmt"
	"os"
	"タイトル/config"
)

// Gitリポジトリのバージョン start.sh実行後バージョン自動更新
var version = "1.0.0"

func CmdFlag() {
	// ポート設定のオプション
	flag.StringVar(&config.FlagPort, "port", config.Config.Port, "ポート設定が可能")
	flag.StringVar(&config.FlagPort, "p", config.Config.Port, "ポート設定が可能(short)")

	// Gitリポジトリのバージョン確認
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "バージョン確認")
	flag.BoolVar(&showVersion, "v", false, "バージョン確認(short)")
	flag.Parse()

	if showVersion {
		// バージョン番号を表示する
		fmt.Println("version", version)
		os.Exit(1)
	}
}

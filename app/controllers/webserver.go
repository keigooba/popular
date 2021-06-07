package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"タイトル/config"

	"github.com/tcnksm/go-latest"
)

func StartWebServer() error {
	files := http.FileServer(http.Dir(config.Config.View))
	http.Handle("/views/", http.StripPrefix("/views/", files))
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "config/config.json") //ファイルにアクセス
	})
	http.HandleFunc("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.FlagPort
		log.Printf("デフォルトのポート %s", port)
	}
	log.Printf("リッスンしているポート %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	// 最新版バージョンチェック
	json := &latest.JSON{
		// JSONを返すURL
		URL: config.Config.URL + fmt.Sprint(config.FlagPort) + "/json",
	}
	res, _ := latest.Check(json, config.Config.Version)
	if res.Outdated {
		fmt.Printf("%sは最新ではない、これはアップグレードすべき %s: %s\n", config.Config.Version, res.Current, res.Meta.Message)
	}
	generateHTML(w, nil, "index")
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf(config.Config.View+"/%s.html", file))
	}
	// ヘッダー・フッターを追加
	files = append(files, config.Config.View+"/_header.html", config.Config.View+"/_footer.html")

	templates := template.Must(template.ParseFiles(files...))
	err := templates.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	controllersTwitter "popular/app/controllers/twitter"
	"popular/config"
	"text/template"

	"github.com/stretchr/objx"
	"github.com/tcnksm/go-latest"
)

var Port = "8080"

func StartWebServer() error {
	files := http.FileServer(http.Dir(config.Config.View))
	http.Handle("/views/", http.StripPrefix("/views/", files))
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "config/config.json") //ファイルにアクセス
	})
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/twitter/oauth", controllersTwitter.TwitterAuthHandler)
	http.HandleFunc("/twitter/callback", controllersTwitter.TwitterCallbackHandler)
	http.HandleFunc("/twitter/post", controllersTwitter.TwitterPostHandler)
	http.HandleFunc("/agreement", commonHandler)
	http.HandleFunc("/privacy_policy", commonHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/", indexHandler)

	Port = os.Getenv("PORT")
	if Port == "" || Port == "8080" {
		Port = config.FlagPort
		log.Printf("デフォルトのポート %s", Port)
	}
	log.Printf("リッスンしているポート %s", Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}

func versionHandler(_ http.ResponseWriter, _ *http.Request) {
	// 最新版バージョンチェック
	json := &latest.JSON{
		// JSONを返すURL
		URL: config.Config.URL + fmt.Sprint(config.FlagPort) + "/json",
	}
	res, _ := latest.Check(json, config.Config.Version)
	if res.Outdated {
		log.Printf("%sは最新ではない、これはアップグレードすべき %s: %s\n", config.Config.Version, res.Current, res.Meta.Message)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "main/index")
	} else {
		generateHTML(w, data, "layout", "private_navbar", "main/index")
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if r.Method == "GET" {
		data, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "main"+url)
		} else {
			generateHTML(w, data, "layout", "private_navbar", "main"+url)
		}
	} else if r.Method == "POST" {
		http.Redirect(w, r, url, http.StatusSeeOther) //キャッシュクリア303指定
	}
}

func commonHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	data, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "main"+url)
	} else {
		generateHTML(w, data, "layout", "private_navbar", "main"+url)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	data, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "main"+url)
	} else {
		data["HomeUser"] = "ログインユーザー" //ナビゲーションメニュー「ホームへ」非表示用
		generateHTML(w, data, "layout", "private_navbar", "main"+url)
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf(config.Config.View+"/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	err := templates.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err)
	}
}

func session(_ http.ResponseWriter, r *http.Request) (data map[string]interface{}, err error) {
	data = map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
		data["PixabayKey"] = config.Config.PixabayKey
	} else {
		return nil, err
	}
	return data, nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
	authCookie.MaxAge = -1 //auth削除
	http.SetCookie(w, authCookie)
	http.Redirect(w, r, "/", 302)
}

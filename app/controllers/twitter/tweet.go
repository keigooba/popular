package controllersTwitter

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"popular/config"

	"github.com/ChimeraCoder/anaconda"
	"github.com/stretchr/objx"
)

// Get ツイートする
func TwitterPostHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("auth")
	if err != nil {
		log.Println("cokkieにデータがありません。")
		http.Redirect(w, r, "/home", 302)
	}
	data := objx.MustFromBase64(authCookie.Value)

	anaconda.SetConsumerKey(config.Config.TwitterConsumerKey)
	anaconda.SetConsumerSecret(config.Config.TwitterConsumerSecret)
	api := anaconda.NewTwitterApi(data["oauth_token"].(string), data["oauth_secret"].(string))

	url_img := r.FormValue("img")
	// 拡張子取得
	e := filepath.Ext(url_img)
	log.Println(e)

	response, err := http.Get(url_img)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	file, err := os.Create("save" + e)
	if err != nil {
		log.Println(err)
	}
	io.Copy(file, response.Body)

	// このままではヘッダーが含まれていない？ので再度開く
	file, _ = os.Open("save" + e)

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ
	file_data := make([]byte, size)
	file.Read(file_data)

	// 画像のアップロード
	base64String := base64.StdEncoding.EncodeToString(file_data)
	media, err := api.UploadMedia(base64String)
	if err != nil {
		log.Println(err)
	}
	// すぐに削除
	if err := os.Remove("save" + e); err != nil {
		log.Println(err)
	}

	v := url.Values{}
	v.Set("media_ids", media.MediaIDString)

	text := r.FormValue("tweet")
	if _, err := api.PostTweet(text, v); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/home", 302)
}

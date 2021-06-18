package controllersTwitter

import (
	"log"
	"net/http"
	"popular/lib/twitter"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/stretchr/objx"
)

// Get ツイートする
func TwitterPostHandler(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("auth")
	if err != nil {
		log.Println("cokkieにデータがありません。")
		http.Redirect(w, r, "/home", 302)
	}
	auth_data := objx.MustFromBase64(authCookie.Value)

	text := r.FormValue("text")
	at := &oauth.Credentials{
		Secret: auth_data["oauth_secret"].(string),
		Token:  auth_data["oauth_token"].(string),
	}
	if err := twitter.PostTweet(at, text); err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/home", 302)
}

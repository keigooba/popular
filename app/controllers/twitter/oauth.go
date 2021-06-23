package controllersTwitter

import (
	"log"
	"net/http"
	"popular/lib/twitter"
)

// Get 認証する
func TwitterAuthHandler(w http.ResponseWriter, r *http.Request) {

	config := twitter.GetConnect()
	callback_url := r.Host
	if callback_url == "localhost:8080" {
		callback_url = "http://localhost:8080/twitter/callback"
	} else {
		callback_url = "https://popular-32pe64nwja-an.a.run.app/twitter/callback"
	}
	rt, err := config.RequestTemporaryCredentials(nil, callback_url, nil)
	if err != nil {
		log.Println(err)
	}

	// セッション保存
	sess := twitter.GlobalSessions.SessionStart(w, r)
	sess.Set("request_token", rt.Token)
	sess.Set("request_token_secret", rt.Secret)

	url := config.AuthorizationURL(rt, nil)

	http.Redirect(w, r, url, 302)
}

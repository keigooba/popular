package controllersTwitter

import (
	"log"
	"net/http"
	"popular/lib/twitter"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/stretchr/objx"
)

// Get コールバックする
func TwitterCallbackHandler(w http.ResponseWriter, r *http.Request) {

	sess := twitter.GlobalSessions.SessionStart(w, r)

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  sess.Get("request_token").(string),
			Secret: sess.Get("request_token_secret").(string),
		},
		r.FormValue("oauth_verifier"),
	)
	if err != nil {
		log.Println(err)
	}

	account := twitter.Account{}
	if err = twitter.GetMe(at, &account); err != nil {
		log.Println(err)
	}
	// 写真が小さいため登録された写真を取得
	imgURL := strings.Replace(account.ProfileImageURL, "_normal", "", -1)

	authCookieValue := objx.New(map[string]interface{}{
		"oauth_secret":      at.Secret,
		"oauth_token":       at.Token,
		"name":              account.ScreenName,
		"avatar_url_origin": imgURL,
		"avatar_url":        account.ProfileImageURL,
	}).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})

	http.Redirect(w, r, "/home", 302)
}

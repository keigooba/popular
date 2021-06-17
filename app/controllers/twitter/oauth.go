package controllersTwitter

import (
	"log"
	"net/http"
	"popular/lib/twitter"

	"github.com/stretchr/objx"
)

// Get 認証する
func TwitterAuthHandler(w http.ResponseWriter, r *http.Request) {

	config := twitter.GetConnect()
	rt, err := config.RequestTemporaryCredentials(nil, "https://popular-32pe64nwja-an.a.run.app/twitter/callback", nil)
	if err != nil {
		log.Println(err)
	}

	authCookieValue := objx.New(map[string]interface{}{
		"request_token":        rt.Token,
		"request_token_secret": rt.Secret,
	}).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})

	url := config.AuthorizationURL(rt, nil)

	http.Redirect(w, r, url, 302)
}

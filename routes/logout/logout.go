package logout

import (
	"fmt"
	"github.com/yuleihua/openid-connect/app"
	"net/http"
	"net/url"
	"os"
	"encoding/json"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	domain := os.Getenv("AUTH0_DOMAIN")
	logoutUrl, err := url.Parse("https://" + domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "/protocol/openid-connect/logout"
	parameters := url.Values{}
	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, _ := json.Marshal(session.Values["profile"])
	fmt.Println("json profile: ", string(content))
	fmt.Println("json profile struct: ", session.Values["profile"])

	token, isOK := session.Values["id_token"]
	if isOK {
		parameters.Add("id_token_hint", token.(string))
	}
	fmt.Printf("token: %v\n", token)

	profile, isOK := session.Values["profile"].(map[string]interface{})
	if isOK {
		parameters.Add("state", profile["session_state"].(string))
	}
	fmt.Printf("state: %v\n", profile)

	parameters.Add("post_logout_redirect_uri", returnTo.String())

	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}

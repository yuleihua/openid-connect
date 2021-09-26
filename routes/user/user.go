package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yuleihua/openid-connect/app"
	"github.com/yuleihua/openid-connect/routes/templates"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, _ := json.Marshal(session.Values["profile"])
	fmt.Printf("\n user profile: %v\n", string(profile))

	state, _ := json.Marshal(session.Values["state"])
	fmt.Printf("\n user state: %v\n", string(state))

	access_token, _ := json.Marshal(session.Values["access_token"])
	fmt.Printf("\n user access_token: %v\n", string(access_token))

	templates.RenderTemplate(w, "user", session.Values["profile"])

	time.Sleep(5 * time.Second)

	http.Redirect(w, r, "/logout", http.StatusSeeOther)
}

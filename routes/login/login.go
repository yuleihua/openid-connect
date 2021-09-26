package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/yuleihua/openid-connect/app"
	"github.com/yuleihua/openid-connect/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Generate random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := app.Store.Get(r, "auth-session")
	fmt.Println("22222session : ", session)
	if err != nil {
		fmt.Printf("session2222 : %v, error: %v\n", session, err)
        session, _ = app.Store.New(r, "auth-session")
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		fmt.Printf("session save, error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("7878778session save : ", session)
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		fmt.Printf("authenticator error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("authenticator : ", state)
	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

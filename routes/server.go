package routes

import (
	"github.com/yuleihua/openid-connect/routes/logout"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/yuleihua/openid-connect/routes/callback"
	"github.com/yuleihua/openid-connect/routes/home"
	"github.com/yuleihua/openid-connect/routes/login"
	"github.com/yuleihua/openid-connect/routes/middlewares"
	"github.com/yuleihua/openid-connect/routes/user"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)

	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	r.HandleFunc("/logout", logout.LogoutHandler)

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}

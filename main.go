package main

import (
	"flag"
	"os"

	"github.com/yuleihua/openid-connect/app"
	"github.com/yuleihua/openid-connect/routes"
)

var (
	domain   string
	client   string
	secret   string
	callback string
)

func init() {
        flag.StringVar(&domain, "domain", "demo-client", "auth2 domain")
	flag.StringVar(&client, "clientID", "localhost/auth/realms/demo", "auth2 client2")
	flag.StringVar(&secret, "secret", "ff603af2-8ee0-4441-96a1-2eee1ac2b5bb", "auth2 secret")
	flag.StringVar(&callback, "callback", "http://localhost:3000/callback", "auth2 callback url")
}

func SetupEnv(isReplace bool) {
	if isReplace {
		os.Setenv("AUTH0_CLIENT_ID", client)
		os.Setenv("AUTH0_DOMAIN", domain)
		os.Setenv("AUTH0_CLIENT_SECRET", secret)
		os.Setenv("AUTH0_CALLBACK_URL", callback)
	}
}

func main() {
	flag.Parse()

	SetupEnv(true)

	app.Init()
	routes.StartServer()
}

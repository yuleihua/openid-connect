package home

import (
	"net/http"

	"github.com/yuleihua/openid-connect/routes/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}

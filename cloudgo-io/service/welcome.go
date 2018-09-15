package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func welcomeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "welcome", struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
		}{
			Username: req.Form["username"][0],
			Password: req.Form["password"][0],
			Email:    req.Form["email"][0],
			Phone:    req.Form["phone"][0],
		})
	}

}

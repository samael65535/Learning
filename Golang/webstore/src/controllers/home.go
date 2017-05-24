package controllers

import (
	"net/http"
	"text/template"
	"controllers/util"
	"models"
	"viewmodels"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	rw := util.GetResponseWriter(w, req)
	defer rw.Close()

	vm := viewmodels.GetLogin()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		member, err := models.GetMember(email, password)
		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId
				rw.Header().Add("Set-Cookie", cookie.String())
			}
		}

	}
	this.loginTemplate.Execute(rw, vm)
}

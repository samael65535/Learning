package controllers

import (
	"encoding/json"
	"net/http"
	"text/template"
	"controllers/util"
	"viewmodels"
)

type standLocatorController struct {
	template *template.Template
}

func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	rw := util.GetResponseWriter(w, req)
	defer rw.Close()

	vm := viewmodels.GetStandLocator()

	rw.Header().Add("Content-Type", "text/html")
	this.template.Execute(rw, vm)

}

func (this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {

	rw := util.GetResponseWriter(w, req)
	defer rw.Close()
	vm := viewmodels.GetStandLocations()
	rw.Header().Add("Content-Type", "application/json")
	data, err := json.Marshal(vm)
	if err == nil {
		rw.Write(data)
	} else {
		rw.WriteHeader(404)
	}
}

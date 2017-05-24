package controllers

import (
	"github.com/gorilla/mux"
	"text/template"
	"net/http"
	"viewmodels"
	"strconv"
)

type productController struct {
	template *template.Template
}

func (this *productController) get (w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err == nil {
		vm := viewmodels.GetProducts(id)
		this.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}

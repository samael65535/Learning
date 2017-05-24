package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
	"controllers/util"
	"converters"
	"models"
	"viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()
	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoryToViewModel(category))
	}

	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)

}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProducts(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	} else {
		w.WriteHeader(404)
	}
}

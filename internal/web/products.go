package web

import (
	"go-storefront/internal/models"
	"html/template"
	"net/http"
)

type ProductsPageData struct {
	Title     string
	CartCount string
	Products  []models.Product
}

func ShowProducts(w http.ResponseWriter, r *http.Request) {
	data := ProductsPageData{
		Title:     "Products",
		CartCount: "0",
		Products: []models.Product{
			{ID: 1, Name: "Nissan", Model: "Leaf", Year: 2013, Price: "5000", ImageURL: "URL"},
			{ID: 2, Name: "BMW", Model: "E60", Year: 2007, Price: "3000", ImageURL: "URL"},
			{ID: 3, Name: "Nissan", Model: "Leaf", Year: 2021, Price: "11500", ImageURL: "URL"},
		},
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/partials/nav.html",
		"templates/products_list.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// base.html should define {{ define "base" }} ... {{ end }}
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

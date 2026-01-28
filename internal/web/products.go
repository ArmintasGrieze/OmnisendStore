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

type ProductPageData struct {
	Title     string
	CartCount string
	Product   models.Product
}

func (h *ProductsHandler) ShowProducts(w http.ResponseWriter, r *http.Request) {
	stmt := "SELECT id, title, description, quantity,  image_url, price FROM products ORDER BY id DESC"
	rows, err := h.DB.Query(stmt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		car := models.Product{}
		err := rows.Scan(&car.ID, &car.Title, &car.Description, &car.Quantity, &car.ImageURL, &car.Price)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		products = append(products, car)

	}

	data := ProductsPageData{
		Title:     "Products",
		CartCount: "0",
		Products:  products,
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

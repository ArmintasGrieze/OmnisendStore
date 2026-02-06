package web

import (
	"go-storefront/internal/models"
	"html/template"
	"net/http"
	"strconv"
)

func (h *ProductsHandler) AdminShowProducts(w http.ResponseWriter, r *http.Request) {
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
		"templates/admin/products_list.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductsHandler) AdminShowEditProduct(w http.ResponseWriter, r *http.Request) {

	// GET - get product's ID from URL
	urlID := r.URL.Query().Get("id")
	if urlID == "" {
		http.Error(w, "Missing ID of the product", http.StatusBadRequest)
	}

	// POST - update product
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		description := r.FormValue("description")
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Please enter quantity in number", http.StatusBadRequest)
			return
		}

		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			http.Error(w, "Please enter quantity in number", http.StatusBadRequest)
			return
		}

		imageUrl := r.FormValue("image_url")

		_, err = h.DB.Exec(`
		UPDATE products
		SET title = ?, description = ?, quantity = ?, price = ?, image_url = ? WHERE id = ?
		`, title, description, quantity, price, imageUrl, urlID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
		return
	}

	// GET products
	car := models.Product{}
	err := h.DB.QueryRow(`
	SELECT id, title, description, quantity,  image_url, price
	FROM products
	WHERE ID = ?`, urlID).Scan(&car.ID, &car.Title, &car.Description, &car.Quantity, &car.ImageURL, &car.Price)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := ProductPageData{
		Title:     "Products",
		CartCount: "0",
		Product:   car,
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/partials/nav.html",
		"templates/admin/edit_product.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *ProductsHandler) AdminShowCreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			http.Error(w, "Please enter quantity in number", http.StatusBadRequest)
			return
		}

		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			http.Error(w, "Please enter quantity in number", http.StatusBadRequest)
			return
		}

		description := r.FormValue("description")
		if err != nil {
			http.Error(w, "Please enter description", http.StatusBadRequest)
			return
		}

		imageUrl := r.FormValue("image_url")
		if err != nil {
			http.Error(w, "Please enter image's URL", http.StatusBadRequest)
			return
		}

		currency := r.FormValue("currency")
		if err != nil {
			http.Error(w, "Please enter currency", http.StatusBadRequest)
			return
		}

		status := r.FormValue("status")
		if err != nil {
			http.Error(w, "Please select product's status", http.StatusBadRequest)
			return
		}

		_, err = h.DB.Exec(`
		INSERT INTO products
		(title, description, quantity, price, image_url, currency, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		`, title, description, quantity, price, imageUrl, currency, status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/partials/nav.html",
		"templates/admin/create_product.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductsHandler) AdminShowDeleteProduct(w http.ResponseWriter, r *http.Request) {

	// GET - get product's ID from URL
	urlID := r.URL.Query().Get("id")
	if urlID == "" {
		http.Error(w, "Missing ID of the product", http.StatusBadRequest)
	}

	_, err := h.DB.Exec(
		`DELETE FROM products WHERE id = ?`,
		urlID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/partials/nav.html",
		"templates/admin/delete_product.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

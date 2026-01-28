package main

import (
	"database/sql"
	"fmt"
	"go-storefront/internal/config"
	"go-storefront/internal/web"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite", cfg.DBPath)
	err = db.Ping()

	h := web.NewProductsHandler(db)

	auth := web.NewAdminAuth(cfg.AdminUser, cfg.AdminPass)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", web.ShowHomepage)
	mux.HandleFunc("/products", h.ShowProducts)
	mux.HandleFunc("/about", web.ShowAbout)
	mux.HandleFunc("/admin/products", auth.Require(h.AdminShowProducts))
	mux.HandleFunc("/admin/products/edit", auth.Require(h.AdminShowEditProduct))
	mux.HandleFunc("/admin/products/new", auth.Require(h.AdminShowCreateProduct))
	mux.HandleFunc("/admin/products/delete", auth.Require(h.AdminShowDeleteProduct))

	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

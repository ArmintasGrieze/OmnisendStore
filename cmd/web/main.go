package main

import (
	"fmt"
	"go-storefront/internal/config"
	"go-storefront/internal/web"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", web.ShowHomepage)
	mux.HandleFunc("/products", web.ShowProducts)
	mux.HandleFunc("/about", web.ShowAbout)

	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		fmt.Println(err)
	}
}

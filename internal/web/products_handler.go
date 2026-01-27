package web

import "database/sql"

type ProductsHandler struct{ DB *sql.DB }

func NewProductsHandler(db *sql.DB) *ProductsHandler { return &ProductsHandler{DB: db} }

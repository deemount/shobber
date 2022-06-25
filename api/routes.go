package api

import (
	"net/http"

	"github.com/deemount/shobber/pkg/home"
	"github.com/deemount/shobber/pkg/products"
)

func Routes() {

	http.HandleFunc("/", home.Handler)
	http.HandleFunc("/products", products.Handler)

}

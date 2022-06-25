package products

import (
	"net/http"

	"github.com/deemount/shobber-tpl/pkg/models"
	"github.com/deemount/shobber-tpl/pkg/view"
	"github.com/deemount/shobber/pkg/meta"
	"github.com/deemount/shobber/pkg/navbar"
)

var tplDir string = "../../lib/ui/html"
var products *view.View

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpls := []string{
		tplDir + "/products.tmpl",
		tplDir + "/pages/products/header.tmpl",
		tplDir + "/pages/products/main.tmpl",
		tplDir + "/pages/products/footer.tmpl",
	}

	data := models.Data{
		PageTitle:       "Shobber:Products",
		PageAuthor:      "Salvatore Gonda",
		PageDescription: "A simple online shop platform",
		OG:              meta.OpenGraph(),
		NavBarLinks:     navbar.NavBar(),
	}

	products = view.NewView(tplDir, "products", tmpls)

	err := products.RenderData(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

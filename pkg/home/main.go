package main

import (
	"net/http"

	"github.com/deemount/shobber-tpl/pkg/models"
	"github.com/deemount/shobber-tpl/pkg/view"
	"github.com/deemount/shobber/pkg/navbar"
)

var tplDir string = "../../lib/ui/html"
var home *view.View

func main() {

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":9090", nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpls := []string{
		tplDir + "/base.tmpl",
		tplDir + "/pages/home/header.tmpl",
		tplDir + "/pages/home/main.tmpl",
		tplDir + "/pages/home/footer.tmpl",
	}

	data := models.Data{
		PageTitle:       "Shobber:Home",
		PageAuthor:      "Salvatore Gonda",
		PageDescription: "A simple online shop platform",
		NavBarLinks:     navbar.NavBar(),
	}

	home = view.NewView(tplDir, "base", tmpls)

	err := home.RenderData(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

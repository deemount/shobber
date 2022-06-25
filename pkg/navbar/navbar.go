package navbar

import "github.com/deemount/shobber-tpl/pkg/models"

func NavBar() []models.Link {

	return []models.Link{
		{Cls: "nav-link px-2 link-secondary", Href: "/", Title: "Home"},
		{Cls: "nav-link px-2 link-dark", Href: "/products", Title: "Products"},
		{Cls: "nav-link px-2 link-dark", Href: "/services", Title: "Services"},
	}

}

package meta

import "github.com/deemount/shobber-tpl/pkg/models"

func OpenGraph() []models.OpenGraph {
	return []models.OpenGraph{
		{Property: "og:title", Content: "Shobber"},
		{Property: "og:type", Content: "ecommerce"},
		{Property: "og:url", Content: "https://www.shobber.de"},
		{Property: "og:description", Content: "A simple online shop platform"},
		{Property: "og:image", Content: "image.png"},
	}
}

package api

import (
	"log"
	"net/http"
)

func Run() {

	go log.Printf("Shobber v%s is ready to listen and serve on port %s", "0.1.0-beta.2", ":9090")

	srv := &http.Server{
		Addr: ":9090",
	}

	log.Fatal(srv.ListenAndServe())

}

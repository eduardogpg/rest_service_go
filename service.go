package main 

/*
	To Use Import

*/

import (
	"net/http"
	"github.com/go-martini/martini"
)

func main() {
	
	m := martini.Classic()
	m.Get("/", func(w http.ResponseWriter, r *http.Request) string {
		return "New REST service"
	})

	
	m.RunOnAddr(":8000")

}


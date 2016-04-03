package main 

/*
	To Use :
	go get github.com/go-sql-driver/mysql
*/

import (
	"net/http"
	"encoding/json"

	"github.com/go-martini/martini"
	
	v "./variables"
	stc "./structures"
	h "./handler"
)

func main() {
	
	h.InitializeUnSafeDB()
	defer h.CloseConnection()

	m := martini.Classic()
	m.Get("/", func(w http.ResponseWriter, r *http.Request) string {
		return "New REST service"
	})

	m.Get("/users/", func(w http.ResponseWriter, r *http.Request) {	
		erros := stc.Errors{}			
		data := h .GetAllUser()

		response := stc.Result { Status : v.Success, Data : data, Errors : erros}
    	json.NewEncoder(w).Encode(response)

	})
	m.RunOnAddr(":8000")

}


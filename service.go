package main 

/*
	To Use :
	go get github.com/go-sql-driver/mysql
*/

import (
	"net/http"
	"encoding/json"
	"strconv"

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
		errors := stc.Errors{}			
		data := h .GetAllUser()

		response := stc.Result { Status : v.Success, Data: data, Errors : errors}
    	json.NewEncoder(w).Encode(response)

	})

	m.Get("/users/:user_id", func(w http.ResponseWriter, r *http.Request, params martini.Params){	
		user_id := params["user_id"]
		errors := stc.Errors{}
		data := stc.Users{}
		
		status, error := IsNumber(user_id)

		if status == v.Success{
			data = h.GetUser(user_id)
			if (len(data) == 0){
        		status = v.Error
        		errors = append(errors, stc.Error{ Body: "User not found"})
    		}
		}else{
			errors = append(errors, error)
		}

		response := stc.Result { Status : status, Data: data ,Errors : errors}
    	json.NewEncoder(w).Encode(response)

	})

	m.RunOnAddr(":8000")
}

func IsNumber(number string) (status int, error stc.Error) {
	_, err := strconv.Atoi(number)
    if err != nil {
        return v.Error, stc.Error{ Body: "The param is not a number"}
    }
	return v.Success, stc.Error{}
}


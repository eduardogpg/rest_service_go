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

	m.Get("/user/:user_id", func(w http.ResponseWriter, r *http.Request, params martini.Params){	
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

	m.Get("/delete/user/:user_id", func(w http.ResponseWriter, r *http.Request, params martini.Params){	
		user_id := params["user_id"]
		errors := stc.Errors{}
		status, error := IsNumber(user_id)
		if status == v.Success{
			if h.DeleteUser(user_id) != true{
				status = v.Error
        		errors = append(errors, stc.Error{ Body: "Is not possible delete the user"})
			}
		}else{
			errors = append(errors, error)
		}
		response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
    	json.NewEncoder(w).Encode(response)

	})

	m.Get("/update/user/:user_id/:user_name", func(w http.ResponseWriter, r *http.Request, params martini.Params){	
		user_id := params["user_id"]
		user_name := params["user_name"]

		errors := stc.Errors{}
		status, error := IsNumber(user_id)
		if status == v.Success{
			if h.UpdateUser(user_id, user_name) != true{
				status = v.Error
        		errors = append(errors, stc.Error{ Body: "Is not possible update the user"})
			}
		}else{
			errors = append(errors, error)
		}
		response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
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


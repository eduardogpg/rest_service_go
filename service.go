package main 

/*
	To Use :
	go get github.com/go-sql-driver/mysql
*/

import (
	connect "./connect"
	handle "./handle"
	"github.com/go-martini/martini"
)

func main() {
	
	connect.InitializeUnSafeDB()
	defer connect.CloseConnection()

	m := martini.Classic()
	m.Get("/users/", handle.GetUsers)
	m.Get("/user/:user_id", handle.GetUser)
	m.Delete("/delete/user/:user_id", handle.DeleteUser)
	m.Put("/update/user/:user_id/:user_name", handle.UpdateUser)
	m.Put("/update/user/", handle.UpdateUserJson)
	m.Post("/new/user/", handle.AddNewUser)
	m.RunOnAddr(":8000")
}

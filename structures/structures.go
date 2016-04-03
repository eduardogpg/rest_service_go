package structures


type User struct {
    User_Id 	string  `json:"user_id"` 
    Created_At  string  `json:"created_at"`
    Update_At   string	`json:"update_at"`
    User_Name 	string  `json:"user_name"`
}

type Users []User


type Error struct {
    Number 	int		`json:"number"`
    Body 	string 	`json:"Body"`

}

type Errors []Error

type Result struct {
	Status 	int			`json:"status"`
	Data 	Users		`json:"data"`
	Errors 	Errors 		`json:"errors"`
}
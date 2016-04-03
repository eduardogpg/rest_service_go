package structures


type User struct {
    User_Id 	string  `json:"user_id"` 
    Created_At  string  `json:"created_at"`
    Update_At   string	`json:"update_at"`
    User_Name 	string  `json:"user_name"`
}

type User []Users
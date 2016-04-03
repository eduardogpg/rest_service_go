package structures


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
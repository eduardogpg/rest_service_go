package handler

import(
	"log"
	"database/sql"
	"fmt"

	stc "../structures"
)

import _ "github.com/go-sql-driver/mysql"

var db *sql.DB
var database string = "test"
var username string = "root"


func InitializeUnSafeDB(){
	db = connect_db(GetConnectionDB(""))
	ChekConnection()
}

func connect_db(connect string) *sql.DB{
    db,err := sql.Open("mysql", connect)
    if(err != nil){
        log.Fatal(err)
        return nil
    }
    return db
}

func GetAllUser() stc.Users{
	allUsers := stc.Users{}
	sentence := "SELECT Id, created_at, update_at, user_name FROM users ORDER BY created_at DESC"
	rows, err := db.Query(sentence)
	log.Println(sentence)

	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) 
        }
        user := stc.User { User_Id: string(values[0]), Created_At: string(values[1]), Update_At : string(values[2]),  User_Name: string(values[3]) }
		allUsers = append(allUsers, user)  
    }
	return allUsers
}

func GetUser(user_id string) stc.Users{
	allUsers := stc.Users{}
	sentence := fmt.Sprintf("SELECT Id, created_at, update_at, user_name FROM users WHERE Id = %s ORDER BY created_at DESC", user_id)
	rows, err := db.Query(sentence)
	log.Println(sentence)

	columns, err := rows.Columns()
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) 
        }
        user := stc.User { User_Id: string(values[0]), Created_At: string(values[1]), Update_At : string(values[2]),  User_Name: string(values[3]) }
		allUsers = append(allUsers, user)  
    }
	return allUsers
}

func ChekConnection(){
	_, err := db.Query("SHOW TABLES")
	if(err != nil){
		log.Fatal(err)
	}
}

func CloseConnection(){
	db.Close()
	log.Println("Connection with Data Base closed")
}

func GetConnectionDB(password string) string{
    return username + ":" + password + "@/" + database
}
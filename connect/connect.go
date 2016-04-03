package connect

import(
	"log"
	"database/sql"
	"fmt"
	"os"
    "time"
    "strings"
    
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

func InitializeSafeDB(environment_password string){
	db = connect_db(GetConnectionDB(os.Getenv(environment_password)))
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

func DeleteUser(user_id string) bool{
    sentence := fmt.Sprintf("DELETE FROM users WHERE Id = %s", user_id)
    _ , error := db.Query(sentence)
    fmt.Println(sentence)
    if error != nil{
    	return false
    }
    return true;
}

func UpdateUser(user_id, user_name string) bool{
    sentence := fmt.Sprintf("UPDATE users SET user_name = '%s' WHERE Id = %s", user_name, user_id)
    _ , error := db.Query(sentence)
    fmt.Println(sentence)
    if error != nil{
    	return false
    }
    return true;
}

func AddUser(user_name string) bool{
    now := GetFormanNow()
    sentence := fmt.Sprintf("INSERT into users (user_name, created_at, update_at) VALUES ('%s', '%s', '%s')", user_name, now, now)
    _ , error := db.Query(sentence)
    fmt.Println(sentence)
    if error != nil{
        return false
    }
    return true;
}

func ChekConnection(){
	_, err := db.Query("SHOW TABLES")
	if(err != nil){
		log.Fatal(err)
	}
}

func GetFormanNow() string{
    return strings.Split(time.Now().String(), " -")[0]
}

func CloseConnection(){
	db.Close()
	log.Println("Connection with Data Base closed")
}

func GetConnectionDB(password string) string{
    return username + ":" + password + "@/" + database
}
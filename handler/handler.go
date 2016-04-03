package handler

import(
	"log"
	"database/sql"
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
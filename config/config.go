package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDB() (db *sql.DB,err error){

	dbName := "trellDb"
	dbUser := "trell"
	dbPass := "<=#trell@rtghvcxdfty@2017#=>"
	dbHost := "trell-mysql-db-staging.cyqwbanzexpw.ap-south-1.rds.amazonaws.com"
	dbPort := "3306"
	url := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?multiStatements=true&parseTime=true"
	//dbName := "trellDb"
	//dbUser := "trell"
	//dbPass := "<=#trell@rtghvcxdfty@2017#=>"
	db, err = sql.Open("mysql", url)
	if err!=nil {
		log.Fatalln("Unable to connect"+string(err.Error()))
	}
	return db,err
}

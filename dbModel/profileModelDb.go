package dbModel

import (
	"database/sql"
	"example/web-service-gin/config"
	"example/web-service-gin/entities"
	"fmt"
	"log"
)

type ProfileModel struct {
Db *sql.DB
}

func (profileModel ProfileModel) FindAll() ([]entities.ProfileUrl,error) {
	db, err :=config.GetDB()
	if err!=nil{
		return nil, err
	}else {
		rows, err2  := db.Query("select * from scraping_ShareChatUrls")
		if err2 !=nil{
			return nil,err2
		}else {
			var profiles []entities.ProfileUrl
			for rows.Next() {
				var profile entities.ProfileUrl
				rows.Scan(&profile.ProfileUrl)
				profiles = append(profiles,profile)
			}
			return profiles,nil
		}
	}
}

func SaveProfileUrl(profileUrl entities.ProfileUrl) error {
	db,e := config.GetDB()
	if e!=nil{
		log.Fatalln(e)
	}
	result, err:= db.Exec("insert into scraping_ShareChatUrls (profileUrls) values (?);",profileUrl.ProfileUrl)
	if err!=nil{
		log.Fatalln(err)
	}else{
		id,_  := result.LastInsertId()
		fmt.Printf(string(id))
		return nil
	}
	return nil
}
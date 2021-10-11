package dbModel

import (
	"database/sql"
	"example/web-service-gin/config"
	"example/web-service-gin/entities"
	"fmt"
	"log"
)

type ProfileDetailModel struct {
	Db *sql.DB
}

func (profileDetailModel ProfileDetailModel) FindAll() ([]entities.ProfileDetails,error){
	db, err :=config.GetDB()
	if err!=nil{
		return nil, err
	}else {
		rows, err2  := db.Query("select * from scraping_profileDetails")
		if err2 !=nil{
			return nil,err2
		}else {
			var profiles []entities.ProfileDetails
			for rows.Next() {
				var profile entities.ProfileDetails
				rows.Scan(&profile.ProfileName,&profile.ProfileHandle,&profile.ProfileIconUrl,&profile.TagLine,&profile.Followers)
				fmt.Printf(profile.ProfileName)
				profiles = append(profiles,profile)
			}
			return profiles,nil
		}
	}
}

func SaveProfileDetails(profileDetails entities.ProfileDetails) error {
	db,e := config.GetDB()
	if e!=nil{
		log.Fatalln(e)
	}
	resultDetails, err:= db.Exec("insert into scraping_profileDetails (profileName, profileHandle, profileIconUrl, TagLine, followers) values (?,?,?,?,?);",
		profileDetails.ProfileName,profileDetails.ProfileHandle,profileDetails.ProfileIconUrl,profileDetails.TagLine,profileDetails.Followers)
	if err!=nil{
		log.Panicln(err)
		return err
	}else{
		id,_  := resultDetails.LastInsertId()
		fmt.Println(string(id))
	}

	for _, dog := range profileDetails.PostUrls {
		resultLinks, err2 := db.Exec("insert into scraping_profileLinks (profileHandle, postLink) values(?,?);",profileDetails.ProfileHandle,dog)
		if err2!=nil {
			log.Panicln(err2)
			return err
		}else{
			resultLinks.RowsAffected()
		}
	}
	return nil
}
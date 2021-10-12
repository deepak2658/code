package main

import (
	"example/web-service-gin/config"
	"example/web-service-gin/dbModel"
	_ "example/web-service-gin/dbModel"
	"example/web-service-gin/entities"
	"example/web-service-gin/kafka"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getProfileDetails(c *gin.Context) {
	db, err := config.GetDB()
	if err !=nil{
		fmt.Println(err)
	}else {
		profileDetailsModel := dbModel.ProfileDetailModel{Db: db}// models.ProductModel{Db:db,}
		ProfileDetails, err := profileDetailsModel.FindAll()
		if err !=nil{
			fmt.Println(err)
		}else {
			c.IndentedJSON(http.StatusOK, ProfileDetails)
		}
	}
}

func postProfileDetails(c *gin.Context) {
	var newProfile entities.ProfileDetails

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	err1 :=dbModel.SaveProfileDetails(newProfile)
	if err1!=nil{
		log.Fatalln("failed to persist data"+err1.Error())
	}

	c.IndentedJSON(http.StatusCreated, newProfile)
}

func main() {
	server := gin.Default()

	server.GET("/all", getProfileDetails)
	server.POST("/add", postProfileDetails)

	server.Run("localhost:8080")
	kafka.StartKafka()
}

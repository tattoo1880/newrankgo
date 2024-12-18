package main

import (
	"gotest/database"
	"gotest/model"
	"gotest/route"
	"log"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

	err := database.MyDB.AutoMigrate(&model.Xhs{}, &model.Weixin{}, &model.Dy{}, &model.Sph{}, &model.Ks{}, &model.Bil{})
	handleErr(err)

	router := route.InitRouter()
	err2 := router.Run("127.0.0.1:8080")
	handleErr(err2)

}

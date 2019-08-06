
package main

import (
	"github.com/gin-gonic/gin"
	"./config"
	"log"
	"fmt"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readJsonConfig() {
	configs, _ := config.ReadConfig("config.json")
	fmt.Println(configs.DB_HOST)
}

func main(){
	
	port := ":8000"

	router := gin.New()
	
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	readJsonConfig()

	router.Run(port)
}
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oGabrielSilva/godailyboost/api/auth"
	"github.com/oGabrielSilva/godailyboost/api/db"
	"github.com/oGabrielSilva/godailyboost/api/routes"
)

func main() {
	auth.ConfigureJWT()

	dbHandler := db.CreateDbHandler()
	defer dbHandler.Client.Close()

	r := gin.Default()

	routes.MapApiRoutesV1(r, dbHandler)

	log.Fatal(r.Run(":8080"))
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

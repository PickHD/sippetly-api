package main

import (
	"log"
	"sippetly-api/db"
	"sippetly-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  	godotenv.Load()
	
	go func() {
		db.Init()
	}()

	r:=gin.Default()
	routes.QuoteRoutes(r)

	if err:=r.Run();err!=nil{
		log.Fatalf("Server Failed with '%s'\n",err)
	}

}

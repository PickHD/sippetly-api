package routes

import (
	"sippetly-api/handlers"

	"github.com/gin-gonic/gin"
)

func QuoteRoutes(q *gin.Engine){
	v1:=q.Group("/api/v1")
	{
		v1.POST("/quotes",handlers.CreateQuoteHandler)
		v1.GET("/quotes",handlers.RetrieveAllQuoteHandler)
		v1.GET("/quotes/:id",handlers.GetOneQuoteHandler)
		v1.PUT("/quotes/:id",handlers.UpdateQuoteHandler)
		v1.DELETE("/quotes/:id",handlers.DeleteQuoteHandler)
	}
}
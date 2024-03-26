package api

import (
	"net/http"

	"github.com/BoburjonRaximov/premier_league/api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(handlar *handler.Handler) *gin.Engine {

	router := gin.Default()
	// add swgger
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hammasi nazoratim ostida")
	})

	baseRouter := router.Group("/api")

	baseRouter.POST("/club", handlar.CreateClub)
	baseRouter.GET("/club/:id", handlar.GetClub)
	baseRouter.PUT("/club/:id", handlar.UpdateClub)
	baseRouter.GET("/club", handlar.GetAllClub)
	baseRouter.DELETE("/club/:id", handlar.DeleteClub)

	return router
}

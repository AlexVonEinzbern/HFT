package main

import (
	"hft/api"
	"hft/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	db.Migrate()
	r := gin.Default()
	p2p := r.Group("/peer2peer")
	{
		peers := p2p.Group("/peers")
		{
			peers.GET(":file", api.GetPeers)
			peers.GET("", api.SearchAllPeers)
			peers.POST("", api.CreatePeer)
		}
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}

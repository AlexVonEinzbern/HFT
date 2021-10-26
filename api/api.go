package api

import (
	"hft/crud"
	"hft/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePeer(c *gin.Context) {
	var cPeer model.PeerCreate
	err := c.BindJSON(&cPeer)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, crud.CreatePeer(cPeer))

}

func SearchAllPeers(c *gin.Context) {
	query := c.Request.URL.Query()
	name := query["name"][0]
	c.IndentedJSON(http.StatusOK, crud.SearchAllPeers(name))
}

func GetPeers(c *gin.Context) {
    query := c.Request.URL.Query()
    file := query["file"][0]
    c.IndentedJSON(http.StatusOK, crud.GetPeers(file))
}

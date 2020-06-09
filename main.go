package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dgkg/videogames/db/mock"
	"github.com/dgkg/videogames/endpoints"
)

var dbMock = mock.New()

func init() {
	mock.MockUser(dbMock.(*mock.DB), "Denis", "Plagnol", "denis@yahoo.fr", "coucou123")
	mock.MockVideoGame(dbMock.(*mock.DB), "Denis", "Plagnol", "denis@yahoo.fr", "coucou123")
}

func main() {
	r := gin.Default()
	endpoints.New(dbMock, r)
	r.Run(":8080")
}

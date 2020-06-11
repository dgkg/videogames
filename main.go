package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dgkg/videogames/db/mock"
	"github.com/dgkg/videogames/endpoints"
)

var dbMock = mock.New()

func init() {
	mock.MockUser(dbMock.(*mock.DB), "Denis", "Plagnol", "denis@yahoo.fr", "46a52db240ac194d7bb1c899d491c835f4de92206ce381fb6ec7a704b1daedb4")
	mock.MockVideoGame(dbMock.(*mock.DB), "Denis", "Plagnol", 1, 1, []string{"coucou123"})
}

func main() {
	r := gin.Default()
	endpoints.New(dbMock, r)
	r.Run(":8080")
}

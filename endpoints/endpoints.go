package endpoints

import (
	"github.com/dgkg/videogames/db"
	"github.com/gin-gonic/gin"
)

// New initialize all endpoints
func New(db db.Store, r *gin.Engine) {
	initEndpointUser(db, r)
	initEndpointVideoGame(db, r)
}

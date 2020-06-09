package endpoints

import (
	"log"
	"net/http"

	"github.com/dgkg/videogames/db"
	"github.com/gin-gonic/gin"
)

type serviceUser struct {
	db db.Store
}

func initEndpointUser(db db.Store, r *gin.Engine) {
	us := &serviceUser{
		db: db,
	}
	r.GET("/users/:uuid", us.Get)
	r.GET("/users", us.GetAll)
	r.DELETE("/users/:uuid", us.Delete)
}

func (su *serviceUser) Get(ctx *gin.Context) {
	u, err := su.db.GetUser(ctx.Param("uuid"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, u)
}

func (su *serviceUser) GetAll(ctx *gin.Context) {
	us, err := su.db.GetUsers()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, us)
}

func (su *serviceUser) Delete(ctx *gin.Context) {
	err := su.db.DeleteUser(ctx.Param("uuid"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, nil)
}

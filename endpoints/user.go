package endpoints

import (
	"log"
	"net/http"

	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/models"
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
	r.POST("/users", us.Create)
	r.PATCH("/users/:uuid", us.UpdateUser)
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

func (su *serviceUser) Create(ctx *gin.Context) {
	var u models.User
	if err := ctx.BindJSON(&u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := su.db.AddUser(&u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, u)
}

func (su *serviceUser) UpdateUser(ctx *gin.Context) {
	update := make(map[string]interface{})

	if err := ctx.BindJSON(&update); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if _, err := su.db.UpdateUser(ctx.Param("uuid"), update); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	u, err := su.db.GetUser(ctx.Param("uuid"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, u)
}

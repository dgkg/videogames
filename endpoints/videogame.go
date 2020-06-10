package endpoints

import (
	"log"
	"net/http"

	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/models"
	"github.com/gin-gonic/gin"
)

type serviceVideoGame struct {
	db db.Store
}

func initEndpointVideoGame(db db.Store, r *gin.Engine) {
	us := &serviceVideoGame{
		db: db,
	}
	r.GET("/videogames/:uuid", us.Get)
	r.GET("/videogames", us.GetAll)
	r.DELETE("/videogames/:uuid", us.Delete)
	r.POST("/videogames", us.Create)
}

func (su *serviceVideoGame) Get(ctx *gin.Context) {
	u, err := su.db.GetVideoGame(ctx.Param("uuid"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, u)
}

func (su *serviceVideoGame) GetAll(ctx *gin.Context) {
	us, err := su.db.GetVideoGames()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, us)
}

func (su *serviceVideoGame) Delete(ctx *gin.Context) {
	err := su.db.DeleteVideoGame(ctx.Param("uuid"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, nil)
}

func (su *serviceVideoGame) Create(ctx *gin.Context) {
	var u models.VideoGame
	if err := ctx.BindJSON(&u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := su.db.AddVideoGame(&u); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(200, u)
}

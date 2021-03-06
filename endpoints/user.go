package endpoints

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"

	"github.com/dgkg/videogames/db"
	"github.com/dgkg/videogames/middleware"
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
	r.POST("/login", us.Login)
	users := r.Group("users")

	users.Use(middleware.NewJWT([]byte("secret")))
	users.GET("/:uuid", us.Get)
	users.GET("/", us.GetAll)
	users.DELETE("/:uuid", us.Delete)
	users.POST("/", us.Create)
	users.PATCH("/:uuid", us.UpdateUser)
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

	u, err := su.db.UpdateUser(ctx.Param("uuid"), update)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(200, u)
}

func (su *serviceUser) Login(ctx *gin.Context) {
	type aux struct {
		Email    string `json:"email"`
		Password string `json:"pass"`
	}

	var a aux
	if err := ctx.BindJSON(&a); err != nil {
		log.Println((err))
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	u, err := su.db.GetUserByEmail(a.Email)

	fmt.Println("u")
	fmt.Println(u)
	if err != nil {
		log.Println((err))
	}

	auxSum := sha256.Sum256([]byte(a.Password))
	if fmt.Sprintf("%x", auxSum) != u.Password {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := middleware.Create([]byte("secret"), u.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
	}
	ctx.JSON(200, token)
}

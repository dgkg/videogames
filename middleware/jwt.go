package middleware

import (
	"net/http"
	"strings"

	"github.com/cristalhq/jwt/v3"
	"github.com/gin-gonic/gin"
)

type JWTPayload struct {
	JWT string `json:"jwt"`
}

func NewJWT(key []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := GetClaim(ctx, key)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}

func Create(key []byte, uuid string) (*JWTPayload, error) {

	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		return nil, err
	}

	claims := &jwt.StandardClaims{
		ID: uuid,
	}

	builder, err := jwt.NewBuilder(signer).Build(claims)
	if err != nil {
		return nil, err
	}

	return &JWTPayload{
		JWT: builder.String(),
	}, nil

}

func GetClaim(ctx *gin.Context, key []byte) (*jwt.Token, error) {
	bearer := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(bearer) > 1 {
		return nil, nil
	}

	token := bearer[1]

	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		return nil, err
	}

	return jwt.ParseAndVerifyString(token, verifier)
}

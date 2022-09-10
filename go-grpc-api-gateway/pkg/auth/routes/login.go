package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/pb"
)

// swagger:model
type LoginResponseBody struct {
	// Error if any
	// example:
	Error string `json:"error,omitempty"`
	// Return status code if any
	// example: 200
	Status int `json:"status"`
	// JWT Authentication token
	// example: token.header.info
	Token string `json:"token"`
}

// swagger:model
type LoginRequestBody struct {
	// the Email for this user
	// required: true
	// example: meet@me.soni
	// swagger:strfmt email
	Email string `json:"email"`
	// the Password for this user
	// required: true
	// min: 8
	// max: 15
	// example: meet@1234
	// swagger:strfmt password
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	body := LoginRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// parameters:
	// - in: body
	//   name: login
	//   description: Login Request Body
	//   schema:
	//      $ref: '#/definitions/LoginRequestBody'

	resp, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	log.Println(resp, err)

	ctx.JSON(http.StatusOK, &resp)
}

package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/pb"
)

// swagger:model
type RegisterRequestBody struct {
	// the Email for this user
	// required: true
	Email string `json:"email"`
	// the Password for this user
	// required: true
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(resp.Status), &resp)
}

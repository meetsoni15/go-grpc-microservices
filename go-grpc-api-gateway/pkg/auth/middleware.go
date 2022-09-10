package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func NewAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (m *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	log.Println("AuthRequired")
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := m.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.User.UserId)
	ctx.Set("user", res.User)

	ctx.Next()

}

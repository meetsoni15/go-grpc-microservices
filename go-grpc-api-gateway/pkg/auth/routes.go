package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth").Use(CORSMiddleware())
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	r.Use(CORSMiddleware())

	return svc
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		log.Println("Here")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

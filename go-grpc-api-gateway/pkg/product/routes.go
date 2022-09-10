package product

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/config"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	authMiddleware := auth.NewAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api")
	{
		routes.POST("/product", svc.CreateProduct)
	}
	routes.Use(CORSMiddleware())
	routes.Use(authMiddleware.AuthRequired)

}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func CORSMiddleware() gin.HandlerFunc {
	log.Println("HITTTT")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		log.Println("asdasdasda")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

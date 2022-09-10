package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meetsoni1511/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc/status"
)

// swagger:model
type CreateProductRequestBody struct {
	//Product name
	//required: true
	// example: Macbook pro M1 Pro
	Name string `json:"name"`
	// Quantity of product that will be inventory
	//required: true
	// example: 13
	Stock int64 `json:"stock"`
	// Price of particular product
	//required: true
	// example: 229000
	Price int64 `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resp, err := c.CreateProduct(ctx.Request.Context(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			// GRPC Error
			ctx.AbortWithError(http.StatusBadGateway, fmt.Errorf(e.Message()))
		} else {
			// Non GRPC Error
			ctx.AbortWithError(http.StatusBadGateway, err)
		}
		return
	}

	ctx.JSON(int(resp.Status), nil)
}

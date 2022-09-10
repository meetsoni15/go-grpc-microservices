package doc

import "github.com/meetsoni1511/go-grpc-api-gateway/pkg/product/routes"

/*
	PRODUCT
*/

// swagger:parameters get-thing update-thing delete-thing
type _ struct {
	// Create product body
	// in: body
	Body routes.CreateProductRequestBody
}

// swagger:operation POST /api/product Auth get-thing
//
// Create Product
// ---
// summary: Create Product
// operationId: createProduct
// produces:
// - application/json
// parameters:
// - in: body
//   name: login
//   description: Login Request Body
//   schema:
//      $ref: '#/definitions/CreateProductRequestBody'
// responses:
//   401:
//     description: Invalid Username || Password
//   404:
//     description: User not found
//   500:
//     description: Internal Server Error

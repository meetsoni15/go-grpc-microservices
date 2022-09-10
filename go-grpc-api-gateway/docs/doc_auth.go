package doc

import "github.com/meetsoni1511/go-grpc-api-gateway/pkg/auth/routes"

/*
	LOGIN
*/

// swagger:parameters userLogin
type _ struct {
	// The body
	// in: body
	Body routes.LoginRequestBody
}

// swagger:operation POST /auth/login Auth userLogin
//
// User Login
// ---
// summary: User login
// operationId: userLogin
// produces:
// - application/json
// responses:
//   200:
//     description: successful operation
//     schema:
//       $ref: '#/definitions/LoginResponseBody'
//   401:
//     description: Invalid Username || Password
//   404:
//     description: User not found
//   500:
//     description: Internal Server Error

/*
	REGISTER
*/

//swagger:parameters userRegister
type _ struct {
	// Request Body
	// in:body
	Body routes.RegisterRequestBody
}

// swagger:operation POST /auth/register Auth userRegister
//
// User Registration
// ---
// summary: User register
// operationId: userRegister
// produces:
// - application/json
// parameters:
// - in: body
//   name: register
//   description: Register Request Body
//   schema:
//      $ref: '#/definitions/RegisterRequestBody'
// responses:
//   201:
//     description: successful operation
//     schema:
//       $ref: '#/definitions/RegisterRequestBody'
//   400:
//     description: Invalid id supplied
//   404:
//     description: User not found

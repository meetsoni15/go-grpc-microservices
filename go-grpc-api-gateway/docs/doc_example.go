package doc

import "time"

// swagger:model ThingResponse
type ThingResponse struct {
	// The UUID of a thing
	// example: 6204037c-30e6-408b-8aaa-dd8219860b4b
	UUID string `json:"uuid"`

	// The Name of a thing
	// example: Some name
	Name string `json:"name"`

	// The Value of a thing
	// example: Some value
	Value string `json:"value"`

	// The last time a thing was updated
	// example: 2021-05-25T00:53:16.535668Z
	Updated time.Time `json:"updated"`

	// The time a thing was created
	// example: 2021-05-25T00:53:16.535668Z
	Created time.Time `json:"created"`
}

// swagger:parameters get-thing update-thing delete-thing
type _ struct {
	// The UUID of a thing
	// in:path
	UUID string `json:"uuid"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// swagger:route GET /thing/{uuidaaa} Thing get-thing
//
// This is the summary for getting a thing by its UUID
//
// This is the description for getting a thing by its UUID. Which can be longer.
//
// responses:
//   200: ThingResponse
//   404: ErrorResponse
//   500: validationError

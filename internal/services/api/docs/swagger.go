// Package classification Website Monitor API
//
// the purpose of this api is to provide a monitoring service for websites
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

// Internal server error message returned as a string
// swagger:response internalServerErrorResponse
type internalServerErrorResponseWrapper struct {
	// error description
	// in: body
	Body MessageError
}

// Not found message error returned as string
// swagger:response notFoundResponse
type errorNotFoundResponseWrapper struct {
	// error description
	// in: body
	Body MessageError
}

// BadRequest message error returned as string
// swagger:response badRequestResponse
type badRequestResponseWrapper struct {
	// error description
	// in: body
	Body MessageError
}

// NoContent message error returned as string
// swagger:response noContentResponse
type noContentResponseWrapper struct { }

type MessageError struct {
	Message string `json:"message"`
}
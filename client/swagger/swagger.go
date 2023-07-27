// Package swagger contains the client library code generated automatically from the swagger definitions.
// The code present in http, models and operations must not be modified.
// To run the source generation, swagger binary must be present in PATH.
// To install the swagger binary, run:
//
//	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
//
// For more information, see: https://github.com/go-swagger/go-swagger
package swagger

//go:generate swagger generate client -f ../../docs/swagger/swagger.json -t ../../ -c client/swagger/http -m client/swagger/models -a /client/swagger/operations
//go:generate go mod tidy

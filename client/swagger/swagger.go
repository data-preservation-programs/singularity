// Package swagger contains the client library code generated automatically from the swagger definitions.
// The code present in http, models and operations must not be modified.
//
// For more information, see: https://github.com/go-swagger/go-swagger
package swagger

//go:generate rm -rf ./http ./models ./operations
//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5 generate client -f ../../docs/swagger/swagger.json -t ../../ -c client/swagger/http -m client/swagger/models -a /client/swagger/operations

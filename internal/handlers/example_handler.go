package handlers

import (
	"github.com/pArtour/networking-api/internal/controllers"
	"github.com/pArtour/networking-api/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func GetExampleHandler() operations.GetExampleHandlerFunc {
	controller := &controllers.ExampleController{}

	return operations.GetExampleHandlerFunc(func(params operations.GetExampleParams) middleware.Responder {
		response, err := controller.GetExample()
		if err != nil {
			return operations.NewGetExampleDefault(500).WithPayload("Internal server error")
		}

		return operations.NewGetExampleOK().WithPayload(response)
	})
}

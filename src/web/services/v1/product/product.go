package product

import (
	_ "github.com/gufranmirza/microservice-proto/proto/v1/error/v1error" //used in swag
	"github.com/gufranmirza/provider-microservice/logging"

	_ "github.com/gufranmirza/provider-microservice/web/renderers" // swag
)

type product struct {
	logger logging.Logger
}

// NewProduct returns product impl
func NewProduct() Product {
	return &product{
		logger: logging.NewLogger(),
	}
}

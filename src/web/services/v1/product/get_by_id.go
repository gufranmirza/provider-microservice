package product

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	_ "github.com/gufranmirza/microservice-proto/proto/v1/error/v1error" //used in swag
	"github.com/gufranmirza/microservice-proto/proto/v1/product/v1product"
	"github.com/gufranmirza/provider-microservice/web/renderers"
)

var products = &v1product.Products{
	Products: []*v1product.Product{
		{
			Id:           "1000",
			Name:         "Dell Ultrasharp Monitor",
			Description:  "Dell 24 inch (60.9 cm) Widescreen Ultrasharp LED Backlit Computer Monitor - Full HD, IPS Panel with VGA, DVI, Display, USB Ports - U2412M (Black)",
			Manufacturer: "Dell",
			Price:        "$200",
			InStock:      true,
			Category:     3,
		},
		{
			Id:           "1001",
			Name:         "LG 27 inch 4K-UHD Monitor",
			Description:  "LG 27 inch 4K-UHD (3840 x 2160) HDR 10 Monitor (Gaming & Design) with IPS Panel, HDMI x 2, Display Port, AMD Freesync  - 27UL500 (Silver Stand with White Body)",
			Manufacturer: "LG",
			Price:        "$300",
			InStock:      true,
			Category:     3,
		},
	},
}

// @Summary Get a Product
// @Description It allows to reterive a Product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param productID path string true "Product ID"
// @Success 200 {object} v1product.Product{}
// @Failure 400 {object} v1error.ErrorResponse{}
// @Failure 404 {object} v1error.ErrorResponse{}
// @Failure 500 {object} v1error.ErrorResponse{}
// @Router /products/productID [GET]
func (p *product) GetByID(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	productID := chi.URLParam(r, "productID")

	for _, p := range products.Products {
		if productID == p.Id {
			render.JSON(w, r, p)
			return
		}
	}

	p.logger.Info(txID).Infof("No product found with the given productID %v", productID)
	render.Render(w, r, renderers.ErrorNotFound(fmt.Errorf("No product found with the given productID %v", productID)))
	return
}

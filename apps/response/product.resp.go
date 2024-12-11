package response

import "synapsis-online-store/apps/entity"

type ProductListResponse struct {
	Id        int    `json:"id"`
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Categorie string `json:"categorie"`
	Stock     int16  `json:"stock"`
	Price     int    `json:"price"`
}

func NewProductListResponseFromEntity(products []entity.Product) []ProductListResponse {
	var productList = []ProductListResponse{}

	for _, product := range products {

		productList = append(productList, ProductListResponse{
			Id:        product.Id,
			SKU:       product.SKU,
			Name:      product.Name,
			Stock:     product.Stock,
			Price:     product.Price,
			Categorie: product.Categorie,
		})
	}

	return productList
}

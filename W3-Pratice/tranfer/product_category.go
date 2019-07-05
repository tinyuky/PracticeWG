package tranfer

import (
	"product-service/entities"
	"product-service/dto"
)

func MapProductForResponse(
	products *[]entities.Products, 
	categoryMap map[int]dto.Category,
)(result []dto.Product, err error){
	for _, product := range *products {
		productDto := dto.Product{
			ID:       product.ID,
			Name:     product.Name,
			Image:    product.Image,
			Category: categoryMap[product.Category],
		}
		result = append(result, productDto)
	}

	return result, nil
}
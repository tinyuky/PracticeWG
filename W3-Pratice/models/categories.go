package models

import (
	"encoding/json"
	"fmt"
	"log"
	"product-service/dto"
	"product-service/entities"
	"product-service/services/categorie_service"
	"strconv"
	"strings"

	"github.com/chuongdang/golang-libs/db"
)

func GetListCategory(idList []string) (categories []entities.Categories, err error) {
	query := "SELECT id_category, name, image FROM category "

	if idList[len(idList)-1] != "" {
		idListString := strings.Join(idList, ",")
		query += fmt.Sprintf("WHERE id_category IN (%v)", idListString)
	}

	err = db.GetConn().Select(&categories, query)

	return
}

func GetCategoryByIdList(products *[]entities.Products) (categoryMap map[int]dto.Category, err error) {
	categoryMap = make(map[int]dto.Category)
	var categoryIdList []string
	for _, product := range *products {
		categoryIdList = append(categoryIdList, strconv.Itoa(product.Category))
	}

	var categoryList []dto.Category
	categoryData, _ := categorie_service.FetchCategoryByIdList(&categoryIdList)
	err = json.Unmarshal(categoryData, &categoryList)

	if err != nil {
		log.Fatal(err)
		return
	}
	for _, category := range categoryList {
		if err != nil {
			return nil, err
		}
		categoryMap[int(category.ID)] = category
	}

	return
}

package models

import (
	"fmt"
	"product-service/entities"
	"strconv"
	"strings"

	"github.com/chuongdang/golang-libs/db"
)

func GetProductsList(page string, limit string, sort string, by string, filter *[]string) (products []entities.Products, err error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt
	query := "SELECT id_product, fk_category, name, image FROM product "

	if filter != nil {
		var filStrArr []string

		for _, val := range *filter {
			filPrsArr := strings.Split(val, ",")
			if len(filPrsArr) == 3 {
				filStr := fmt.Sprintf("%v %v '%v'", filPrsArr[0], filPrsArr[1], filPrsArr[2])
				filStrArr = append(filStrArr, filStr)
			}
		}
		if len(filStrArr) > 0 {
			query += "WHERE " + strings.Join(filStrArr, " AND ")
		}
	}

	query += fmt.Sprintf(" ORDER BY %v %v LIMIT %v, %v ", sort, by, offset, limit)
	err = db.GetConn().Select(&products, query)

	return
}

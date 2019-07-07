package categorie_service

import (
	"io/ioutil"
	"net/http"
	"product-service/config"
	"product-service/helper"
	"strings"
)

const PATH_CATEGORY_LIST = "categories"

func FetchCategoryByIdList(idList *[]string) (response []byte, err error){
	req, err := http.NewRequest("GET", config.CATEGORY_SERVICE_ENDPOINT + PATH_CATEGORY_LIST, nil)

	if err != nil {
		return
	}
	q := req.URL.Query()
	uniIDList := helper.UniStrSlice(idList)
	q.Add("ids", strings.Join(uniIDList, ","))

	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return
}
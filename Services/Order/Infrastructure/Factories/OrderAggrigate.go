package order_factories

import (
	"bytes"
	order_enums "delivery/Services/Order/Application/Enums"
	order_domain_aggrigate "delivery/Services/Order/Domain/Aggrigates"
	order_infrastructues_models "delivery/Services/Order/Infrastructure/Models"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func NewOrderAggrigate(CostomerId int, prods map[int]int) (*order_domain_aggrigate.OrderAggrigate, error) {

	ids := ""
	for i, _ := range prods {
		id, err := strconv.ParseInt(i, 10, 32)
		if err != nil {
			continue
		}

		ids += fmt.Sprintf("%d,", id)
	}

	ids = strings.TrimSuffix(ids, ",")

	products, err := getProducts(ids)
	if err != nil {
		return nil, err
	}

	cost := float32(0)
	for _, product := range products {
		pQty, _ := prods[product.Id]
		cost += float32(pQty) * product.Price
	}
	uuid := uuid.New()
	Order := order_infrastructues_models.OrderModel{
		CostumerId: CostomerId,
		Cost:       cost,
		Status:     order_enums.ORDER_CREATED,
	}
	var Items []product_domain_entities.ProductEntity
	var Price float32
}

func getProducts(ids string) ([]*product_infrastructure_models.Product, error) {

	Client := http.Client{
		Timeout: time.Second * 2,
	}
	data := map[string]string{
		"ids": ids,
	}
	d, _ := json.Marshal(data)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/product/multiple", bytes.NewReader(d))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	dd := shared_models.ResponseModel{}
	err = json.Unmarshal(body, &dd)
	if err != nil {
		return nil, err
	}
	products, ok := dd.Result.(map[string]any)["products"]
	if !ok {
		return nil, errors.New("products are missing in the response")
	}
	product := &product_infrastructure_models.Product{}
	productsResult := []*product_infrastructure_models.Product{}
	for _, p := range products.([]interface{}) {
		mapstructure.Decode(p, product)
		productsResult = append(productsResult, product)
	}
	return productsResult, nil

}

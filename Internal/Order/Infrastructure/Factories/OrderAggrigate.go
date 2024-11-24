package order_factories

import (
	"bytes"
	order_enums "delivery/Internal/Order/Application/Enums"
	order_domain_aggrigate "delivery/Internal/Order/Domain/Aggrigates"
	order_infrastructues_models "delivery/Internal/Order/Infrastructure/Models"
	product_infrastructure_models "delivery/Internal/Product/Infrastructure/Models"
	shared_models "delivery/Internal/Shared/Infrastructure/Models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func NewOrderAggrigate(CostomerId int, prods map[int]int) (*order_domain_aggrigate.OrderAggrigate, error) {

	ids := ""
	for id := range prods {
		ids += fmt.Sprintf("%d,", id)
	}

	ids = strings.TrimSuffix(ids, ",")

	products, err := getProducts(ids)
	if err != nil {
		return nil, err
	}

	cost := float32(0)
	for _, product := range products {
		pQty := prods[product.Id]
		cost += float32(pQty) * product.Price
	}
	fingerPrint := uuid.New().String()
	Order := order_infrastructues_models.Order{
		CostumerId:       CostomerId,
		OrderFingerprint: fingerPrint,
		Cost:             cost,
		Status:           order_enums.ORDER_CREATED,
	}

	return &order_domain_aggrigate.OrderAggrigate{
		Order: &Order,
		Items: products,
		Price: cost,
	}, nil

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

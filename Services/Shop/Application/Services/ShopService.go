package shop_application_services

import (
	shop_domain_contracts "delivery/Services/Shop/Domain/Contracts"
)

type ShopService struct {
	repo shop_domain_contracts.IShopRepository
}

func NewShopService(repo shop_domain_contracts.IShopRepository) *ShopService {
	return &ShopService{
		repo: repo,
	}
}

func (obj *ShopService) OpenShop() error {

	return nil
}

func (obj *ShopService) CloseShop() error {

	return nil
}

func (obj *ShopService) UpdateShopInfo() error {
	return nil
}

func (obj *ShopService)GetShopStatus() (string, error) {
	status := 
}

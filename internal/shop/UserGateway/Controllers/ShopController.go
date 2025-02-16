package shop_usergateway_controllers

import shop_application_services "delivery/internal/shop/Application/Services"

type ShopController struct {
	service *shop_application_services.ShopService
}

package shop_usergateway_controllers

import shop_application_services "delivery/Internal/Shop/Application/Services"

type ShopController struct {
	service *shop_application_services.ShopService
}

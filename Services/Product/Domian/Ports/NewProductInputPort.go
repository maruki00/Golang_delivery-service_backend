package product_domain_ports

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	adapters_viewmodels "delivery/app/adapters/ViewModels"
)

type NewProductOutputPort interface {
	handel(requestmodel product_domain_dtos.ProductDTO) adapters_viewmodels.JsonViewModel
}

package product_services

import product_domain_repositories "delivery/Services/Product/Domian/Repositories"

type ProductServices struct {
	productRepository product_domain_repositories.IProductRepository
}

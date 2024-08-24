package product_domain_repositories

import product_domain_entities "delivery/Services/Product/Domian/Entities"

type IProductRepository interface {
	Insert(product product_domain_entities.ProductEntity) (product_domain_entities.ProductEntity, error)
	GetById(id int) (product_domain_entities.ProductEntity, error)
	Search(seasrch string) ([]product_domain_entities.ProductEntity, error)
	Update(id int, data map[string]interface{}) (product_domain_entities.ProductEntity, error)
	Delete(id int) (product_domain_entities.ProductEntity, error)
}

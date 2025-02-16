package product_domain_repositories

import (
	"context"
	product_domain_entities "delivery/internal/product/Domian/Entities"
	product_infrastructure_models "delivery/internal/product/Infrastructure/Models"
)

type IProductRepository interface {
	Insert(ctx context.Context, product product_domain_entities.ProductEntity) (product_domain_entities.ProductEntity, error)
	GetById(ctx context.Context, id int) (product_domain_entities.ProductEntity, error)
	Search(ctx context.Context, seasrch string) ([]product_infrastructure_models.Product, error)
	Update(ctx context.Context, id int, data map[string]interface{}) (product_domain_entities.ProductEntity, error)
	Delete(ctx context.Context, id int) (product_domain_entities.ProductEntity, error)
	GetProductByMultipleId(ctx context.Context, ids []int) ([]product_infrastructure_models.Product, error)
}

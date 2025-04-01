package products

type ProductServiceInterface interface {
	GetProduct()
	GetAllProduct()
	DeleteProduct()
	CreateProduct()
}

type ProductService struct {
	repo ProductRepository
}

func NewRepoProductService(sql ProductRepository) *ProductService {
	return &ProductService{repo: sql}
}

func (service *ProductService) CreateProduct(product *Product) error {
	return service.repo.Create(product)
}
func (service *ProductService) GetProduct(id int) (Product, error) {
	return service.repo.FindById(id)
}
func (service *ProductService) GetAllProduct() ([]Product, error) {
	return service.repo.FindAll()
}
func (service *ProductService) DeleteProduct(id int) error {
	return service.repo.DeleteById(id)
}

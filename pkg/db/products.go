package db

import (
	"github.com/levelstudio/payroll-4ta-crud/pkg/models"
	"gorm.io/gorm"
)

type ProductsRepo struct {
	DB                    *gorm.DB
	CreateProductObserver map[string]chan *models.Product
	UpdateProductObserver map[string]chan *models.Product
	DeleteProductObserver map[string]chan *models.Product
}

func (m *ProductsRepo) GetProducts() ([]*models.Product, error) {

	var products []*models.Product

	query := m.DB.Model(&models.Product{}).
		Order("id ASC").
		Find(&products)
	if err := query.Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (m *ProductsRepo) GetProduct(id string) (*models.Product, error) {

	var product *models.Product

	query := m.DB.Model(&models.Product{}).
		Where("id = ?", id).
		Find(&product)
	if err := query.Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (m *ProductsRepo) CreateProduct(product *models.Product) (*models.Product, error) {

	query := m.DB.Create(product)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.CreateProductObserver {
		observer <- product
	}

	return product, nil
}

func (m *ProductsRepo) UpdateProduct(product *models.Product) (*models.Product, error) {

	query := m.DB.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(product)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.UpdateProductObserver {
		observer <- product
	}

	return product, nil
}

func (m *ProductsRepo) DeleteProduct(product *models.Product) (*models.Product, error) {

	query := m.DB.Unscoped().
		Delete(&product)
	if err := query.Error; err != nil {
		return nil, err
	}

	for _, observer := range m.DeleteProductObserver {
		observer <- product
	}

	return product, nil
}

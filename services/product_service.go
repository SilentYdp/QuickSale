package services

import (
	"fmt"
	"myprojects/flash_sale/datamodels"
	"myprojects/flash_sale/repositories"
)

type IProductService interface {
	GetProductByID(int64)(*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	DeleteProductByID(int64)bool
	InsertProduct(product *datamodels.Product)(int64,error)
	UpdateProduct(product *datamodels.Product)error
}

type ProductService struct {
	ProductRepository repositories.IProduct
}

func NewProductService(repository repositories.IProduct) IProductService {
	return &ProductService{
		ProductRepository:repository,
	}
}
func (p *ProductService)  GetProductByID(ID int64)(*datamodels.Product,error) {
	product, err := p.ProductRepository.SelectByKey(ID)
	if err!=nil{
		fmt.Println(err.Error())
		return &datamodels.Product{},err
	}
	return product,nil
}

func (p *ProductService) GetAllProduct()([]*datamodels.Product,error) {
	products, err := p.ProductRepository.SelectAll()
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return products,nil
}

func (p *ProductService) DeleteProductByID(ID int64)bool {
	b := p.ProductRepository.Delete(ID)
	return b
}

func (p *ProductService) InsertProduct(product *datamodels.Product)(int64,error) {
	id, err := p.ProductRepository.Insert(product)
	if err!=nil{
		fmt.Println(err.Error())
		return -1,err
	}
	return id,nil
}

func (p *ProductService) UpdateProduct(product *datamodels.Product)error {
	err := p.ProductRepository.Update(product)
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}
	return nil
}
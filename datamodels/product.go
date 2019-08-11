package datamodels

type Product struct {
	ID int64	`json:"id" sql:"ID" ye:"id"`
	ProductName string	`json:"productName" sql:"product_name" ye:"productName"`
	ProductNum int64	`json:"productNum" sql:"product_num" ye:"productNum"`
	ProductImg string	`json:"productImg" sql:"product_img" ye:"productImg"`
	ProductUrl string	`json:"productUrl" sql:"product_url" ye:"productUrl"`
}
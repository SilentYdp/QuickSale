package repositories

import (
	"database/sql"
	"fmt"
	"myprojects/flash_sale/datamodels"
	"strconv"
)

//定义接口
//实现接口

type IProduct interface {
//	数据库连接
	Conn()(error)
	Insert(product *datamodels.Product)(int64,error)
	Delete(ID int64)bool
	Update(product *datamodels.Product)error
	SelectByKey(ID int64)(*datamodels.Product,error)
	SelectAll()([]*datamodels.Product,error)
}

type ProductManager struct {
	Tabel string
	MysqlConn *sql.DB
}

func (p *ProductManager) Conn()(err error) {
	if p.MysqlConn == nil {
		dbConn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/sale?charset=utf8")

		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		p.MysqlConn = dbConn
	}
		if p.Tabel == "" {
			p.Tabel = "product"
		}
	return nil
}

func (p *ProductManager)Insert(product *datamodels.Product)(Id int64,err error)  {
	if err:=p.Conn();err!=nil{
		return
	}
	sql:="INSERT INTO `sale`.`product` (`product_name`, `product_num`, `product_img`, `product_url`) VALUES (?,?,?,?);"
	stmt, err := p.MysqlConn.Prepare(sql)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	result, err:=stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	Id, _= result.LastInsertId()
	return
}

func (p *ProductManager)Delete(Id int64)bool  {
	if err:=p.Conn();err!=nil{
		fmt.Println(err.Error())
		return false
	}

	sql:="delete from `sale`.`product` where ID=?"
	stmt, err := p.MysqlConn.Prepare(sql)
	if err!=nil{
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(Id)
	if err!=nil{
		return false
	}
	return true
}

func (p *ProductManager)Update(product *datamodels.Product)(err error)  {
	if err=p.Conn();err!=nil{
		fmt.Println(err.Error())
		return err
	}

	sql:="UPDATE `sale`.`product` SET `product_name` =?,`product_num`=?,product_img=?,product_url=? WHERE `id` ="+strconv.FormatInt(product.ID,10)

	stmt, err := p.MysqlConn.Prepare(sql)
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImg, product.ProductUrl)
	if err!=nil{
		fmt.Println(err.Error())
		return err
	}
	return
}

func (p* ProductManager)SelectByKey(productId int64)(product_ret *datamodels.Product,err error)  {
	if err=p.Conn();err!=nil{
		return &datamodels.Product{},err
	}
	sql:="select product_name,product_num,product_img,product_url from product where id="+strconv.FormatInt(productId,10)


	stmt, err := p.MysqlConn.Prepare(sql)
	if err!=nil{
		fmt.Println(err.Error())
		return &datamodels.Product{},err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	row.Scan(&product_ret.ProductName,&product_ret.ProductNum,
		&product_ret.ProductImg,&product_ret.ProductUrl)
	return
}

func (p *ProductManager)SelectAll()([]*datamodels.Product, error)  {
	if err:=p.Conn();err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	var arr []*datamodels.Product

	sql:="select product_name,product_num,product_img,product_url from product"

	stmt, err := p.MysqlConn.Prepare(sql)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}
	for rows.Next(){
		var pd *datamodels.Product
		rows.Scan(&pd.ProductName,&pd.ProductNum,&pd.ProductImg,&pd.ProductUrl)
		arr=append(arr,pd)
	}
	return arr,nil
}
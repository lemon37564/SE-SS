package product

import (
	"database/sql"
	"encoding/json"
	"se/database"
)

type Product struct {
	fn *database.ProductDB
}

func ProductInit(db *sql.DB) *Product {
	p := new(Product)
	p.fn = database.ProductDBInit(db)
	return p
}

func (p Product) AddProduct(pdname string, price int, description string, amount int, account string, bid bool, date string) bool {
	_, err := p.fn.AddNewProduct(pdname, price, description, amount, account, bid, date)
	if err != nil {
		return false
	}
	return true
}

func (p *Product) DeleteProduct(pdid int) bool {
	err := p.fn.Delete(pdid)
	if err != nil {
		return false
	}

	return true
}
func (p *Product) ChangePrice(pdid, price int) string {
	err := p.fn.UpdatePrice(pdid, price)
	if err != nil {
		return "Price cannot change"
	}
	return "Price has been changed"
}

func (p *Product) ChangeAmount(pdid, amount int) string {

	err := p.fn.UpdateAmount(pdid, amount)
	if err != nil {
		return "Amount cannot change"
	}
	return "Amount change success"
}

func (p *Product) ChangeDescription(pdid int, description string) string {

	err := p.fn.UpdateDescription(pdid, description)
	if err != nil {
		return "Description cannot change"
	}
	return "Description change success"
}

func (p *Product) SetEvaluation(pdid int, eval float64) string {
	err := p.fn.UpdateEval(pdid, eval)
	if err != nil {
		return "Evaluation cannot change"
	}
	return "Evaluation change success"
}

// SearchProductsByName return products info in json
func (p *Product) SearchProductsByName(name string) string {
	pds := p.fn.Search(name)

	res, err := json.Marshal(pds)
	if err != nil {
		panic(err)
	}

	return string(res)
}

// debugging only
func (p *Product) GetAll() string {
	pds := p.fn.All()

	res, err := json.Marshal(pds)
	if err != nil {
		panic(err)
	}

	return string(res)
}

func (p *Product) GetProductInfo(uid int) (products string) {
	//var orders string = ""
	return p.fn.GetInfoFromPdID(uid).StringForProduct()
}

func (p *Product) GetProdPrice(pdid int) int { //拿價格
	return p.fn.GetInfoFromPdID(pdid).Price
}

func (p *Product) GetProAmount(pdid int) int { //拿數量
	return p.fn.GetInfoFromPdID(pdid).Amount
}

func (p *Product) GetProdDescription(pdid int) string { //拿說明
	return p.fn.GetInfoFromPdID(pdid).Description
}

func (p *Product) GetProdDate(pdid int) string { //商品釋出日期
	return p.fn.GetInfoFromPdID(pdid).Date
}

func (p *Product) GetProdName(pdid int) string { //拿商品名稱
	return p.fn.GetInfoFromPdID(pdid).PdName
}

func (p *Product) GetProdEval(pdid int) float64 { //拿評價
	return p.fn.GetInfoFromPdID(pdid).Eval
}

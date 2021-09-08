package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	s := json.NewEncoder(w)
	return s.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	s := json.NewDecoder(r)
	return s.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getProductlastID()
	productList = append(productList, p)

}

func getProductlastID() int32 {
	id := int32(len(productList)) - int32(1)
	return id + 1
}

var productList = []*Product{
	{
		ID:    1,
		Name:  "Samsung",
		Price: "$200",
	},
	{
		ID:    2,
		Name:  "Nokia",
		Price: "$900",
	},
}

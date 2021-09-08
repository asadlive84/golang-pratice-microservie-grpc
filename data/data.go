package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int32   `json:"id" validate:"gt=0,id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	validate.RegisterValidation("id", validateID)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matched := re.FindAllString(fl.Field().String(), -1)
	if len(matched) != 1 {
		return false
	}
	return true
}

func validateID(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`\d+`)
	matched := re.FindAllString(fl.Field().String(), -1)
	if len(matched) != 1 {
		return false
	}
	return true
}

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	s := json.NewEncoder(w)
	return s.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getProductlastID()
	productList = append(productList, p)

}

func UpdateProduct(id int32, p *Product) error {
	log.Println("=============> product", p)
	_, i, err := findProductByID(id)

	if err != nil {
		return err
	}
	p.ID = id
	productList[i] = p
	log.Printf("=============> product %+v", productList[i])
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProductByID(id int32) (*Product, int32, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, int32(i), nil
		}
	}
	return nil, 0, ErrorProductNotFound
}

func getProductlastID() int32 {
	product := productList[int32(len((productList))-1)]
	return product.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Samsung",
		Description: "Taking mobile",
		Price:       200,
	},
	{
		ID:          2,
		Name:        "Nokia",
		Description: "Taking mobile",
		Price:       900,
	},
}

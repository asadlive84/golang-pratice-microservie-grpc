package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		ID:    1,
		Name:  "ujjj",
		Price: 2500,
		SKU:   "sd-gh-jk",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

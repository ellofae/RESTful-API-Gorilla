package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(prod *Product) {
	prod.ID = getProductID()
	productList = append(productList, prod)
}

func UpdateData(id int, prod *Product) error {
	pos, err := getProductPosition(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[pos] = prod

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product was not found")

func getProductPosition(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getProductID() int {
	p := productList[len(productList)-1]
	return p.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Title:       "Chocolate cake",
		Description: "A fluffy cake made of Alpine dark chocolate",
		Price:       5.99,
		SKU:         "soag214f",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Title:       "Brownie",
		Description: "A tasty chocolate make with berries",
		Price:       1.99,
		SKU:         "fas412a",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Title:       "Croissant",
		Description: "A crunchy delisious make of bread and vanilla",
		Price:       2.99,
		SKU:         "opf123h",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

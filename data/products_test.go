package data

import (
	"testing"
)

func TestValidateStructTitle(t *testing.T) {
	t.Run("Title validation test one", func(t *testing.T) {
		prod := &Product{
			Title:       "test",
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Title validation test two", func(t *testing.T) {
		prod := &Product{
			Title:       "test123",
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("Regex is not correct for the title field context")
		}
	})

	t.Run("Title validation test three", func(t *testing.T) {
		prod := &Product{
			Title:       "412@34#-_",
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("Regex is not correct for the title field context")
		}
	})

	t.Run("Title validation test three", func(t *testing.T) {
		prod := &Product{
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("The field is supposed to be with context value")
		}
	})
}

func TestValidateStructDescription(t *testing.T) {
	t.Run("Description validation description one", func(t *testing.T) {
		prod := &Product{
			Title:       "test",
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Description validation test two", func(t *testing.T) {
		prod := &Product{
			Title: "test",
			Price: 1.99,
			SKU:   "abs-abs-abs",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("Regex is not correct for the title field context")
		}
	})
}

func TestValidateStructSKU(t *testing.T) {
	t.Run("SKU validation description one", func(t *testing.T) {
		prod := &Product{
			Title:       "test",
			Description: "test description for the product",
			Price:       1.99,
			SKU:         "abs-abs-abs",
		}
		err := prod.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("SKU validation test two", func(t *testing.T) {
		prod := &Product{
			Title: "test",
			Price: 1.99,
			SKU:   "abs-abs",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("Regex is not correct for the title field context")
		}
	})

	t.Run("SKU validation test three", func(t *testing.T) {
		prod := &Product{
			Title: "test",
			Price: 1.99,
			SKU:   "123-452-bas",
		}
		err := prod.Validate()

		if err == nil {
			t.Fatal("Regex is not correct for the title field context")
		}
	})
}

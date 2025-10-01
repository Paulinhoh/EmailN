package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			println("Validation failed on field:", v.StructField(), "with tag:", v.Tag())
		}
	}
}
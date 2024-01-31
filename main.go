package main

import (
	"github.com/evandersondev/emailn/internal/domain/campaign"
	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Not be error")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField()+" is required with min:", v.Param())
			case "max":
				println(v.StructField()+" is required with max:", v.Param())
			case "email":
				println(v.StructField() + " is invalid")
			}
		}
	}
}

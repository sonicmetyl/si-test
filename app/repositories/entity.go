package repositories

import "gopkg.in/go-playground/validator.v9"

type (
	Taxes struct {
		ID      string  `json:"-"`
		TaxName string  `json:"tax_name", valid:"required"`
		TaxCode string  `json:"tax_code", valid:"required,min:1,max:3"`
		Amount  float64 `json:"amount", valid:"required,numeric"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}

	TaxCode struct {
		ID   string `json:"code_id"`
		Name string `json:"code_name"`
	}

	CalculatedTaxes struct {
		TaxName     string  `json:"tax_name"`
		Amount      float64 `json:"amount"`
		TaxCode     int64   `json:"tax_code"`
		TaxType     string  `json:"tax_type"`
		TaxAmount   float64 `json:"tax_amount"`
		TotalAmount float64 `json:"total_amount"`
	}
)
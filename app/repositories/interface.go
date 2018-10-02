package repositories

import "github.com/labstack/echo"

type Reader interface{
	GetBillings() (billings []CalculatedTaxes, TotalAmount, TotalTaxAmount, FinalTotalAmount float64, err error)
	GetTaxCodes() (taxCodes []TaxCode, err error)
	ValidateTax(c echo.Context) (tax Taxes, err error)
	ValidateTaxCode(taxCode string) bool
}

type Writer interface{
	InsertTax(tax *Taxes) (err error)
}

type StorageInterface interface {
	Reader
	Writer
}
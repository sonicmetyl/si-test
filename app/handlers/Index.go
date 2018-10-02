package handlers

import (
	"github.com/labstack/echo"
	"go-sql/app/adapters"
	"go-sql/app/repositories/mysql"
	"net/http"
	"go-sql/app/repositories"
	"go-sql/app/helpers"
)

type SI struct {
	Storage repositories.StorageInterface
}

func GetIndex(c echo.Context) error {
	test := mysql.NewStorage(adapters.DbReadConn, adapters.DbWriteConn)

	billings, TotalAmount, TotalTaxAmount, FinalTotalAmount, err := test.GetBillings()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"billings": billings,
			"totalAmount" : TotalAmount,
			"finalTaxAmount" : TotalTaxAmount,
			"finalTotalAmount" : FinalTotalAmount,
			"createEndPoint": helpers.GetConfigString("BASE_URL") + "/new",
		})
}

func NewTax(c echo.Context) error {
	repo := mysql.NewStorage(adapters.DbReadConn, adapters.DbWriteConn)
	taxCodes, err := repo.GetTaxCodes()
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "new.html", map[string]interface{}{
		"taxCodes": taxCodes,
		"postUrl" : helpers.GetConfigString("BASE_URL") + "/new",
	})
}

func CreateTax(c echo.Context) error {
	repo := mysql.NewStorage(adapters.DbReadConn, adapters.DbWriteConn)
	tax, err:= 	repo.ValidateTax(c)
	if err != nil {
		return err
	}

	if err := repo.InsertTax(&tax); err != nil {
		return c.Redirect(http.StatusMovedPermanently, helpers.GetConfigString("BASE_URL") + "/new")
	}
	return c.Redirect(http.StatusMovedPermanently, helpers.GetConfigString("BASE_URL"))
}
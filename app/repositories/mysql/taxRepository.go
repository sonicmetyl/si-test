package mysql

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"go-sql/app/repositories"
	"strconv"
)

func (storage Storage) GetBillings() (billings []repositories.CalculatedTaxes, TotalAmount, TotalTaxAmount, FinalTotalAmount float64, err error) {
	sqlString := `
		SELECT t.tax_name, t.tax_code, tc.name as tax_type, t.amount
		FROM taxes t
		INNER JOIN tax_codes tc ON tc.id = t.tax_code
	`
	rows, err := storage.readConn.Raw(sqlString).Rows()
	if err != nil {
		return nil, 0, 0, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var tax repositories.CalculatedTaxes
		if err := rows.Scan(&tax.TaxName, &tax.TaxCode, &tax.TaxType, &tax.Amount); err != nil {
			return billings, 0,0, 0, errors.Wrap(err, "failed to scan order GetBillings")
		}
		switch tax.TaxCode {
			case 1: //
				tax.TaxAmount = 0.1 * tax.Amount
				break
			case 2:
				tax.TaxAmount = 10 + (0.02 * tax.Amount)
				break
			case 3:
				if tax.Amount < 100 {
					tax.TaxAmount = 0
				} else {
					tax.TaxAmount = 0.01 * (tax.Amount - 100)
				}
				break
			default:
				tax.TaxAmount = 0 // default set to zero
		}

		tax.TotalAmount = tax.TaxAmount + tax.Amount
		billings = append(billings, tax)
		TotalTaxAmount = TotalTaxAmount + tax.TaxAmount
		TotalAmount = TotalAmount + tax.Amount
		FinalTotalAmount = FinalTotalAmount + tax.TotalAmount
	}
	return billings,TotalAmount,TotalTaxAmount,FinalTotalAmount , nil
}

func (storage Storage) GetTaxCodes() (taxCodes []repositories.TaxCode, err error) {
	sqlString := `SELECT id, name FROM tax_codes`
	rows, err := storage.readConn.Raw(sqlString).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var code repositories.TaxCode
		if err := rows.Scan(&code.ID, &code.Name); err != nil {
			return nil, errors.Wrap(err, "failed to scan taxCode object")
		}
		taxCodes = append(taxCodes, code)
	}
	return taxCodes, nil
}

func (storage Storage) InsertTax(tax *repositories.Taxes) (err error) {
	var db = storage.writeConn.DB()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to start transaction")
	}

	queryString := `INSERT INTO taxes(tax_name, tax_code, amount) values (?,?,?)`
	stmt, err := tx.Prepare(queryString)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to create statement")

	}
	defer stmt.Close()

	if _, err := stmt.Exec(tax.TaxName, tax.TaxCode, tax.Amount); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to write transaction")
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}

func (storage Storage) ValidateTax(c echo.Context) (tax repositories.Taxes, err error) {
	TaxName := c.FormValue("tax_name")
	if TaxName == "" {
		return tax, errors.Wrap(errors.New("tax name"), "should not empty")
	}
	TaxCode := c.FormValue("tax_code")
	if TaxCode == "" {
		return tax, errors.Wrap(errors.New("tax code"), "should not empty")
	}
	if exist := storage.ValidateTaxCode(TaxCode); !exist {
		return tax, errors.Wrap(errors.New("tax code"), "invalid value")
	}
	Amount := c.FormValue("amount")
	if TaxCode == "" {
		return tax, errors.Wrap(errors.New("tax amount"), "should not empty")
	}
	fAmount , err := strconv.ParseFloat(Amount, 10)
	if err != nil {
		return tax, errors.Wrap(errors.New("tax amount"), "should be numeric")
	}
	tax.TaxName = TaxName
	tax.TaxCode = TaxCode
	tax.Amount  = fAmount

	return tax, nil
}

func (storage Storage) ValidateTaxCode(taxCode string) bool {
	var taxId string
	sqlString := `SELECT id FROM tax_codes where id=?`
	storage.readConn.Raw(sqlString, taxCode).Row().Scan(&taxId)
	if taxId != "" {
		return true
	}
	return false
}
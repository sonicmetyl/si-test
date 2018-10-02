package repositories

import (
	"github.com/pkg/errors"
	. "go-sql/app/adapters"
)

//InitMigration
func InitMigration() (err error) {
	var db = DbWriteConn.DB()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to start migration transaction")
	}

	dbCreate := `CREATE DATABASE IF NOT EXISTS si_test`
	stmt, err := tx.Prepare(dbCreate)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to create DB")

	}
	defer stmt.Close()

	if _, err := stmt.Exec(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to write transaction")
	}

	taxCodeTable := `
		CREATE TABLE IF NOT EXISTS tax_codes (
  			id int(11) NOT NULL AUTO_INCREMENT,
  			name varchar(20) DEFAULT NULL,
  			PRIMARY KEY (id),
			UNIQUE KEY UniqueCode (name)
		) ENGINE=InnoDB;
	`
	stmtTaxCode, err := tx.Prepare(taxCodeTable)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to create tax_codes table ")
	}
	defer stmtTaxCode.Close()
	if _, err := stmtTaxCode.Exec(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to write transaction create table tax_codes")
	}

	// INSERT TaxCode

	insertCodes := `INSERT IGNORE INTO tax_codes VALUES (1,'food'),(2,'tobacco'),(3,'entertainment')`
	stmtTaxData, err := tx.Prepare(insertCodes)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to insert tax_codes ")
	}
	defer stmtTaxData.Close()
	if _, err := stmtTaxData.Exec(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to write transaction insert tax_codes")
	}

	// Create Table Taxes
	taxTable := `
		CREATE TABLE IF NOT EXISTS taxes (
  			id int(11) unsigned NOT NULL AUTO_INCREMENT,
  			tax_name varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  			tax_code int(11) DEFAULT NULL,
  			amount decimal(10,2) DEFAULT NULL,
		PRIMARY KEY (id),
  		KEY tax_code_ref_idx (tax_code),
  		CONSTRAINT tax_code_ref FOREIGN KEY (tax_code) REFERENCES tax_codes (id)
		) ENGINE=InnoDB
	`
	stmtTaxTable, err := tx.Prepare(taxTable)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to create table taxes")
	}
	defer stmtTaxTable.Close()
	if _, err := stmtTaxTable.Exec(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to write transaction create taxes table")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}
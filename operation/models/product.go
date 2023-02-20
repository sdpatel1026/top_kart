package models

import (
	"database/sql"
)

type Product struct {
	ID             int64  `json:"product_id"`
	Name           string `binding:"required" json:"product_name"`
	ActualPrice    int64  `binding:"required" json:"actual_price"`
	FinalPrice     int64  `binding:"required" json:"final_price"`
	TotalUnits     int64  `json:"total_units"`
	AvailableUnits int64  `json:"available_units"`
	ExpiryTime     string `binding:"required" json:"expiry_time"`
}

func (mySQL MySQL) CreateProduct(product Product) (int64, error) {
	insertQuery := "INSERT into products(product_name,actual_price,final_price,total_units,available_units,expiry_time)values(?,?,?,?,?,?)"
	tx, err := mySQL.DB.Begin()
	if err != nil {
		return -1, err
	}
	result, err := tx.Exec(insertQuery, product.Name, product.ActualPrice, product.FinalPrice, product.TotalUnits, product.AvailableUnits, product.ExpiryTime)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	productID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	return productID, err

}

func (mySQL MySQL) UpdateProduct(product Product) error {
	updateQuery := "UPDATE products SET product_name=?,actual_price=?,final_price=?,total_units=?,available_units=?,expiry_time=? WHERE id=?"
	_, err := mySQL.DB.Exec(updateQuery, product.Name, product.ActualPrice, product.FinalPrice, product.TotalUnits, product.AvailableUnits, product.ExpiryTime, product.ID)
	return err
}

func (tx *Tx) GetAvailableProductCount(productID int64) (int64, error) {
	var availableProductCount sql.NullInt64

	selelctQuery := "SELECT available_units FROM products WHERE id=?"
	row := tx.tx.QueryRow(selelctQuery, productID)
	err := row.Scan(&availableProductCount)
	return availableProductCount.Int64, err
}

func (tx *Tx) updateAvailableProductCount(productID, availableUnits int64) error {
	updateQuery := "UPDATE products SET available_units=? WHERE id=?"
	_, err := tx.tx.Exec(updateQuery, availableUnits, productID)
	return err

}

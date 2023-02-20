package models

import (
	"log"
	"time"
)

type Product struct {
	ID             int64  `json:"product_id"`
	Name           string `json:"product_name"`
	ActualPrice    int64  `json:"actual_price"`
	FinalPrice     int64  `json:"final_price"`
	TotalUnits     int64  `json:"total_units"`
	AvailableUnits int64  `json:"available_units"`
	ExpiryTime     string `json:"expiry_time"`
}

func (tx *Tx) getProduct(productID int64) (Product, error) {
	selectQuery := `SELECT * FROM products WHERE id=?`
	var product Product
	row := tx.tx.QueryRow(selectQuery, productID)
	err := row.Scan(&product.ID, &product.Name, &product.ActualPrice, &product.FinalPrice, &product.TotalUnits, &product.AvailableUnits, &product.ExpiryTime)
	return product, err
}

func (mySql MySQL) GetUnExpiredProductDeal() ([]Product, error) {
	selectQuery := "select * FROM products where available_units>0"

	products := make([]Product, 0)
	rows, err := mySql.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	currentTime, _ := time.Parse("15:04:05", time.Now().UTC().Format("15:04:05"))
	for rows.Next() {
		var product Product
		var times []uint8
		err = rows.Scan(&product.ID, &product.Name, &product.ActualPrice, &product.FinalPrice, &product.TotalUnits, &product.ActualPrice, &times)
		if err != nil {
			log.Printf("error in scaning product from the db:%s\n", err.Error())
			continue
		}
		expiryTime, _ := time.Parse("15:04:05", string(times))
		if currentTime.Before(expiryTime) {
			product.ExpiryTime = string(times)
			products = append(products, product)
		}

	}

	return products, err
}

package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Error string
type Tx struct {
	tx *sql.Tx
}
type MySQL struct {
	DB *sql.DB
}

var mySQL MySQL

const (
	DEAL_EXPIRED = "deal expired"
)

func GetDb() MySQL {
	return mySQL
}
func Connect(user string, password string, dbHostAddr string, dbName string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dbHostAddr, dbName)
	var err error
	mySQL.DB, err = sql.Open("mysql", connectionString)
	return err
}

func (mySql MySQL) CreateOrder(order *Order) (err error) {
	insertQuery := "INSERT into orders(product_id,no_of_units,status,actual_cost,final_cost)values(?,?,?,?,?)"
	var tx Tx

	tx.tx, err = mySql.DB.Begin()
	if err != nil {
		return err
	}
	product, err := tx.getProduct(order.ProductID)
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	currentTime, _ := time.Parse("15:04:05", time.Now().UTC().Format("15:04:05"))
	expiryTime, _ := time.Parse("15:04:05", product.ExpiryTime)
	if currentTime.After(expiryTime) {
		tx.tx.Rollback()
		return errors.New(DEAL_EXPIRED)
	}
	order.FinalCost = order.NumberOfUnits * product.FinalPrice
	order.ActualCost = order.NumberOfUnits * product.ActualPrice
	result, err := tx.tx.Exec(insertQuery, order.ProductID, order.NumberOfUnits, ORDER_PLACED, order.ActualCost, order.FinalCost)
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	order.ID, err = result.LastInsertId()
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	return tx.tx.Commit()
}

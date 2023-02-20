package models

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Error string
type Tx struct {
	tx *sql.Tx
}
type MySQL struct {
	DB *sql.DB
}

const (
	STATUS_REJECTES = "status_rejected"
	ITEM_SORTAGE    = "Sortage of items"
)

var mySQL MySQL

func GetDb() MySQL {
	return mySQL
}
func Connect(user string, password string, dbHostAddr string, dbName string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dbHostAddr, dbName)
	var err error
	mySQL.DB, err = sql.Open("mysql", connectionString)
	return err
}

func (mySQL MySQL) ApproveOrder(orderID int) error {
	var tx Tx
	var err error
	tx.tx, err = mySQL.DB.Begin()
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	order, err := tx.GetOrder(int64(orderID))
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	if order.STATUS == int16(OREDER_REJECTED) {
		tx.tx.Rollback()
		return errors.New(STATUS_REJECTES)
	} else if order.STATUS == int16(ORDER_APPROVED) {
		tx.tx.Rollback()
		return nil
	}
	availableItems, err := tx.GetAvailableProductCount(order.ProductID)
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	if availableItems-order.NumberOfUnits < 0 {
		tx.tx.Rollback()
		return errors.New(ITEM_SORTAGE)
	}
	err = tx.updateAvailableProductCount(order.ProductID, availableItems-order.NumberOfUnits)
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	err = tx.approveOrder(orderID)
	if err != nil {
		tx.tx.Rollback()
		return err
	}
	if availableItems-order.NumberOfUnits == 0 {
		err = tx.rejectAllOrdersForProduct(orderID, order.ProductID)
		if err != nil {
			tx.tx.Rollback()
			return err
		}
	}
	err = tx.tx.Commit()
	return err
}

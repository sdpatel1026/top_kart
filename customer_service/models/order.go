package models

import "database/sql"

type STATUS int16
type Order struct {
	ID            int64 `json:"order_id"`
	ProductID     int64 `json:"product_id" binding:"required" `
	NumberOfUnits int64 `json:"no_of_units"`
	STATUS        int16 `json:"status"`
	ActualCost    int64 `json:"actual_cost"`
	FinalCost     int64 `json:"final_cost"`
}

const (
	ORDER_PLACED STATUS = iota
	ORDER_APPROVED
	OREDER_REJECTED
)

func (mySql MySQL) GetOrderStatus(orderID int64) (int16, error) {
	selectQuery := `SELECT status FROM orders WHERE id=?`
	row := mySql.DB.QueryRow(selectQuery, orderID)
	var status sql.NullInt16
	err := row.Scan(&status)
	return status.Int16, err
}

package models

type STATUS int16
type Order struct {
	ID            int64 `json:"order_id"`
	ProductID     int64 `binding:"required" json:"product_id"`
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

func (tx *Tx) GetOrder(orderID int64) (*Order, error) {
	var order Order
	selectQuery := `SELECT product_id,no_of_units,status FROM orders WHERE id=?`
	row := tx.tx.QueryRow(selectQuery, orderID)
	err := row.Scan(&order.ProductID, &order.NumberOfUnits, &order.STATUS)
	return &order, err
}

func (tx *Tx) approveOrder(orderID int) error {

	updateQuery := "UPDATE orders SET status=? WHERE id=?"
	_, err := tx.tx.Exec(updateQuery, ORDER_APPROVED, orderID)
	return err
}

func (tx *Tx) rejectAllOrdersForProduct(orderID int, productID int64) error {

	updateQuery := "UPDATE orders SET status=? WHERE id !=? AND product_id=? AND status=?"

	_, err := tx.tx.Exec(updateQuery, OREDER_REJECTED, orderID, productID, ORDER_PLACED)
	return err
}

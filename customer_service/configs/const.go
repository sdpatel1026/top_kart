package configs

const (
	KEY_MSG    string = "msg"
	KEY_STATUS string = "status"

	STATUS_ERROR   int = 0
	STATUS_SUCCESS int = 1

	INVALID_ORDER_ID   string = "Invalid order_id."
	INVALID_PRODUCT_ID string = "Invalid product_id."
	DEAL_EXPIRED       string = "Deal that you are looking for is expired, please try after 00:00 UTC"
	INVALID_INPUT      string = "Invalid input."
	TECHNICAL_ERROR    string = "Some technical error occured, please try after some time..."
	NO_DEAL_PRESENT    string = "No deal present at this moment."
)

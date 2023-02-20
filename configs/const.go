package configs

const (
	KEY_MSG    string = "msg"
	KEY_STATUS string = "status"

	STATUS_ERROR   int = 0
	STATUS_SUCCESS int = 1

	INVALID_INPUT                string = "Invalid input."
	INVALID_ORDER_ID             string = "Invalid order_id."
	INVALID_PRODUCT_ID           string = "Invalid product_id."
	DEAL_EXPIRED                 string = "Deal that you are looking for is expired, please try after 00:00 UTC"
	TECHNICAL_ERROR              string = "Some technical error occured, please try after some time..."
	NO_DEAL_PRESENT              string = "No deal present at this moment."
	PRODUCT_SUCCESSFULLY_UPDATED string = "Product successfully updated."
	OREDER_REJECTED              string = "Order already rejected."
	ORDER_SUCCESSFULLY_APPROVED  string = "Order successfully approved."
	NO_SUFFICIENT_ITEM           string = "No sufficient amounts of item present."
	EXPIRTY_TIME_EXCEEDED        string = "Expiry time should not be more than 12:00:00"
	INVALID_EXPIRY_TIME          string = "Invalid expiry time."
	REGION_MISSING               string = "region required."
	INVALID_REGION               string = "region must be from (EUR, AR, AFR, AMR)"
	ORDER_ID_MISSING             string = "order_id required."
)

var REGIONS map[string]bool = map[string]bool{"AR": true, "AFR": true, "AMR": true, "EUR": true}

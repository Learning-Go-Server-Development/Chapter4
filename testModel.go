package goproxy

type Order struct {
	ID          int64    `json:"id"`
	CID         int64    `json:"cid"`
	OrderNumber string   `json:"oid"`
	Items       *[]int64 `json:"items"`
}

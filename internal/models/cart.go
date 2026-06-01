package models

type CartItem struct {
	ID           uint    `json:"item_id" gorm:"primarykey"`
	UserID       uint    `json:"-"`
	MedicineID   uint    `json:"medicine_id"`
	Quantity     int     `json:"quantity"`
	PricePerUnit int     `json:"price_per_unit"`
	LineTotal    float64 `json:"line_total"`
}

type CartUpsert struct {
	MedicineID uint `json:"medicine_id"`
	Quantity   int  `json:"quantity"`
}

type Cart struct { 
	UserID     uint    `json:"user_id"`
	Items      []CartItem  `json:"items"`
	TotalPrice float64 `json:"total_price"`
}

type CartUpdateQuantity struct {
	Quantity int `json:"quantity"`
}

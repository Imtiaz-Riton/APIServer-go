package entity

type Product struct {
	ID 	string  `json:"id"`
	Title 	string  `json:"title"`
	Amount 	int64   `json:"amount"`
	Price 	float64 `json:"price"`
}

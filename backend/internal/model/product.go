package model

type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

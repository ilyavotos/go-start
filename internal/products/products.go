package products

type Product struct {
	Id      int    `json:"id,omitempty"`
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}

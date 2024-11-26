package expenses

// Expenses represents the output structure.
type Expenses struct {
	Items []struct {
		Name  string  `json:"name"`
		Price float32 `json:"price"`
	} `json:"items"`
	Total    float32 `json:"total"`
	Category string  `json:"category"`
}

package expenses

// Expenses represents the output structure.
type Expenses struct {
	Items    []Item  `json:"items"`
	Total    float32 `json:"total"`
	Category string  `json:"category"`
}

type Item struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

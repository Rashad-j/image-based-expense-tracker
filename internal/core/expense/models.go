package expenses

// Expenses represents the output structure.
type Expenses struct {
	Items []struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	} `json:"items"`
	Total string `json:"total"`
}

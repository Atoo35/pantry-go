package types

type GetPantryResponse struct {
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Errors        []interface{} `json:"errors"`
	Notifications bool          `json:"notifications"`
	PercentFull   int           `json:"percentFull"`
	Baskets       []struct {
		Name string `json:"name"`
		TTL  int    `json:"ttl"`
	} `json:"baskets"`
}

type UpdatePantryDetailsInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

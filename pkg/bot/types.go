package bot

type APIResponse struct {
	Success bool            `json:"success"`
	Map     *map[string]any `json:"map"`
	ListMap []Product       `json:"listMap"`
}

type Product struct {
	IsActive   string `json:"is_active"`
	ProductURL string `json:"product_url"`
	Price      string `json:"price"`
	FeSKU      string `json:"fe_sku"`
	Locale     string `json:"locale"`
}

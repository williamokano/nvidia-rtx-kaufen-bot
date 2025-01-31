package bot

type APIResponse struct {
	Success bool            `json:"success"`
	Map     *map[string]any `json:"map"`
	ListMap []any           `json:"listMap"`
}

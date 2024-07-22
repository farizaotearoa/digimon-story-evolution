package request

type DigimonListRequest struct {
	PageSize  int      `json:"pageSize"`
	PageNum   int      `json:"pageNum"`
	SortBy    string   `json:"sort_by"`
	SortOrder string   `json:"sort_order"`
	Stage     []string `json:"stage"`
	Type      []string `json:"type"`
	Attribute []string `json:"attribute"`
}

package presenter

type Tag struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListTagResp struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Results []*Tag  `json:"results"`
}

type TagResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Results *Tag   `json:"results"`
}

package tag

type TagPayload struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TagUpdatePayload struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TagDeletePayload struct {
	Id int `json:"id"`
}

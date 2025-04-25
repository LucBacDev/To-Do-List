package todo

type TodoPayload struct {
	UserId      int `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"` 
	DueDate     string `json:"due_date"` 
}

type TodoUpdatePayload struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
}

type TodoDeletePayload struct {
	Id int `json:"id"`
}
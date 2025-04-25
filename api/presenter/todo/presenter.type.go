package presenter

type Todo struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListTodoResp struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Results []*Todo `json:"results"`
}

type TodoResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Results *Todo  `json:"results"`
}

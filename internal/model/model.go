package model

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"created_at"`
}

type Comment struct {
	Id        string `json:"id"`
	PostId    string `json:"post_id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

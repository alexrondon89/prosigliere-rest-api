package model

import "time"

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	Id        string    `json:"id"`
	PostId    string    `json:"blog_post_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

package model

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	IsRead bool   `json:"is_read"`
}

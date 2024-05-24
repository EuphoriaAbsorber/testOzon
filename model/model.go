package model

type Post struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Text              string `json:"text"`
	AuthorID          int    `json:"authorid"`
	IsCommentsUnabled bool   `json:"iscommentsunabled"`
}

type Comment struct {
	ID           int    `json:"id"`
	PostID       int    `json:"postid"`
	ParentCommID int    `json:"parentcommid"`
	AuthorID     int    `json:"authorid"`
	Text         string `json:"text"`
}

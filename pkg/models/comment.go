//pkg/models/comment.go

package models

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Body   string `json:"body"`
}
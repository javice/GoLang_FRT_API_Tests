// pkg/models/album.go

package models

type Album struct {
	UserID 	int `json:"userID"`
	ID		int `json:"id"`
	Title	string `json:"title"`
}
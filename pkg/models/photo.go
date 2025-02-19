// pkg/models/photo.go

package models

type Photo struct {
	AlbumID 	int `json:"albumID"`
	ID		int `json:"id"`
	Title	string `json:"title"`
	URL		string `json:"url"`
	ThumbnailURL	string	`json:"thumbnailUrl"`

}
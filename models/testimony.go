package models

type Testimony struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Photo         string `json:"photo"`
	Content       string `json:"content"`
	Name          string `json:"name"`
	DestinationID uint   `json:"destination_id"`
}

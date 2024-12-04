package models

type Destination struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Photo    string     `json:"photo"`
	Location string     `json:"location"`
	Price    float64    `json:"price"`
	Reviews  []Testimony `json:"reviews" gorm:"foreignKey:DestinationID"`
}

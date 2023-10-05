package data

import "time"

type LeatherGoods struct {
	ID          int64     // Unique integer ID for the leather good
	CreatedAt   time.Time // Timestamp for when the leather good is added to our database
	Name        string    // Name of the leather good
	Type        string    // Type of the leather good
	Price       float64   // Price of the leather good
	LeatherType string    // Type of leather (e.g., natural, eco-friendly)
	Color       string    // Color of the leather good
	Version     int32     // The version number starts at 1 and will be incremented each
	// time the leather good information is updated
}

package models

// Book contains info about books in library
type Book struct {
	ID      int     `json:"id" bson:"_id"`
	Name    string  `json:"name" bson:"name"`
	Price   float32 `json:"price" `
	Edition string  `json:"edition"`
}

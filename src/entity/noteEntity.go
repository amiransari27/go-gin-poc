package entity

type Note struct {
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}

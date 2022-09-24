package models

type Message struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Content   string `json:"content" bson:"content"`
	SendTime  int64  `json:"send_time" bson:"send_time"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

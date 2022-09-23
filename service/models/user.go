package models

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type UserResponse struct {
}
type UserRequest struct {
	ID string
}

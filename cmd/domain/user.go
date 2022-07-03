package domain

// UserRequest represents user request type
type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// UserResponse represents user response type
type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// User represent user type
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserModel type
type UserModel struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

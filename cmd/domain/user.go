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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

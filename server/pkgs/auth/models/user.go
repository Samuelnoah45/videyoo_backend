package auth_models

type User struct {
	ID          string   `json:"id"`
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name" validate:"required"`
	Email       string   `json:"email" validate:"required,min=6,email"`
	Password    string   `json:"password"`
	UserRoles   []string `json:"user_roles"`
	PhotoUrl    string   `json:"photo_url"`
	Gender      string   `json:"gender"`
	PhoneNumber string   `json:"phone_number" validate:"required,min=6"`
}

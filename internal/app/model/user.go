package model

type User struct {
	ID              string `json:"user_id"`
	RoleID          string `json:"role_id"`
	FullName        string `json:"full_name"`
	ProfilePicture  string `json:"profile_picture"`
	Description     string `json:"description"`
	PhoneNumber     string `json:"phone_number"`
	PersonalWebsite string `json:"personal_website"`
	IsBuyer         bool   `json:"is_buyer"`
}

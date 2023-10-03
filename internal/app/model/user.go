package model

type User struct {
	ID              string `json:"user_id" firestore:"user_id"`
	RoleID          string `json:"role_id" firestore:"role_id"`
	FullName        string `json:"full_name" firestore:"full_name"`
	ProfilePicture  string `json:"profile_picture" firestore:"profile_picture"`
	Description     string `json:"description" firestore:"description"`
	PhoneNumber     string `json:"phone_number" firestore:"phone_number"`
	PersonalWebsite string `json:"personal_website" firestore:"personal_website"`
	IsBuyer         bool   `json:"is_buyer" firestore:"is_buyer"`
}

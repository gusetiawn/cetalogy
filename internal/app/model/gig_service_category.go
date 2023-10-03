package model

type GigServiceCategory struct {
	ID            string `json:"id" firestore:"id"`
	ParentID      int    `json:"parent_id" firestore:"parent_id"`
	Url           string `json:"url" firestore:"url"`
	ImageFileName string `json:"image_file_name" firestore:"image_file_name"`
	Description   string `json:"description" firestore:"description"`
	IsActive      bool   `json:"is_active" firestore:"is_active"`
	CategoryName  string `json:"category_name" firestore:"category_name"`
}

package model

type GigServiceCategory struct {
	ID            string `json:"id"`
	ParentID      int    `json:"parent_id"`
	Url           string `json:"url"`
	ImageFileName string `json:"image_file_name"`
	Description   string `json:"description"`
	IsActive      bool   `json:"is_active"`
	CategoryName  string `json:"category_name"`
}

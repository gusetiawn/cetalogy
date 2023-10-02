package model

type GigService struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	CategoryID     string `json:"category_id"`
	Title          string `json:"title"`
	GigServiceTags string `json:"gig_service_tags"`
	Descriptions   string `json:"descriptions"`
	Faq            string `json:"faq"`
}

package model

type GigService struct {
	ID             string `json:"id" firestore:"id"`
	UserID         string `json:"user_id" firestore:"user_id"`
	CategoryID     string `json:"category_id" firestore:"category_id"`
	Title          string `json:"title" firestore:"title"`
	GigServiceTags string `json:"gig_service_tags" firestore:"gig_service_tags"`
	Descriptions   string `json:"descriptions" firestore:"descriptions"`
	Faq            string `json:"faq" firestore:"faq"`
}

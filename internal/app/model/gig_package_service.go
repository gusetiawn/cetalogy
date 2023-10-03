package model

type GigPackageService struct {
	ID           string `json:"id" firestore:"id"`
	PackageID    string `json:"package_id" firestore:"package_id"`
	ServiceID    string `json:"service_id" firestore:"service_id"`
	Price        string `json:"price" firestore:"price"`
	Delivery     string `json:"delivery" firestore:"delivery"`
	DeliveryUnit string `json:"delivery_unit" firestore:"delivery_unit"`
}

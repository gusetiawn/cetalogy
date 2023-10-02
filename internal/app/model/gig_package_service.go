package model

type GigPackageService struct {
	ID           string `json:"id"`
	PackageID    string `json:"package_id"`
	ServiceID    string `json:"service_id"`
	Price        string `json:"price"`
	Delivery     string `json:"delivery"`
	DeliveryUnit string `json:"delivery_unit"`
}

package model

type Package struct {
	ID   string `json:"id" firestore:"id"`
	Name string `json:"name" firestore:"name"`
}

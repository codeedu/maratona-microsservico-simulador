package entity

type Order struct {
	Uuid        string `json:"order"`
	Destination string `json:"destination"`
}

type Destination struct {
	Order string `json:"order"`
	Lat   string `json:"lat"`
	Lng   string `json:"lng"`
}

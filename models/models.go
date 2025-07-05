package models

type Shipment struct {
	ID          int     `json:"id"`
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Weight      float64 `json:"weight"`
	Urgency     string  `json:"urgency"`
	CreatedAt   string  `json:"created_at"`
}

type ShippingOption struct {
	Provider      string  `json:"provider"`
	EstimatedTime int     `json:"estimated_time"`
	Cost          float64 `json:"cost"`
	Score         float64 `json:"score"`
}

package model

type SocialData struct {
	EcoId       string
	Source      string
	ExternalId  []string
	Data        interface{}
	AutoMapping bool
	Version     float64
	UpdateTime  int64
	Country     string
}



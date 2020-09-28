package gateways

import "time"

// Gateway ...
type Gateway struct {
	Token                 string        `json:"token"`
	GatewayType           string        `json:"gateway_type"`
	Description           interface{}   `json:"description"`
	PaymentMethods        []string      `json:"payment_methods"`
	State                 string        `json:"state"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	Name                  string        `json:"name"`
	Characteristics       []string      `json:"characteristics"`
	Credentials           []interface{} `json:"credentials"`
	GatewaySettings       struct{}      `json:"gateway_settings"`
	GatewaySpecificFields []string      `json:"gateway_specific_fields"`
	Redacted              bool          `json:"redacted"`
	Sandbox               bool          `json:"sandbox"`
}

// GatewayJSON ...
type GatewayJSON struct {
	Gateway `json:"gateway"`
}

// GatewayList ...
type GatewayList struct {
	Gateways []Gateway `json:"gateways"`
}

// NewGateway represents the data required to create a new Gateway
// type NewGateway struct {
// 	Token       string
// 	GatewayType string `json:"gateway_type"`
// }

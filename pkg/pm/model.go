package pm

import "time"

type PaymentMethod struct {
	Token        string    `json:"token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Data         struct{}  `json:"data"`
	StorageState string    `json:"storage_state"`
	Test         bool      `json:"test"`
	Metadata     struct {
		Key        string `json:"key"`
		AnotherKey int    `json:"another_key"`
		FinalKey   bool   `json:"final_key"`
	} `json:"metadata"`
	LastFourDigits         string        `json:"last_four_digits"`
	FirstSixDigits         string        `json:"first_six_digits"`
	CardType               string        `json:"card_type"`
	FirstName              string        `json:"first_name"`
	LastName               string        `json:"last_name"`
	Month                  int           `json:"month"`
	Year                   int           `json:"year"`
	Address1               interface{}   `json:"address1"`
	Address2               interface{}   `json:"address2"`
	City                   interface{}   `json:"city"`
	State                  interface{}   `json:"state"`
	Zip                    interface{}   `json:"zip"`
	Country                interface{}   `json:"country"`
	PhoneNumber            interface{}   `json:"phone_number"`
	Company                interface{}   `json:"company"`
	FullName               string        `json:"full_name"`
	EligibleForCardUpdater bool          `json:"eligible_for_card_updater"`
	ShippingAddress1       interface{}   `json:"shipping_address1"`
	ShippingAddress2       interface{}   `json:"shipping_address2"`
	ShippingCity           interface{}   `json:"shipping_city"`
	ShippingState          interface{}   `json:"shipping_state"`
	ShippingZip            interface{}   `json:"shipping_zip"`
	ShippingCountry        interface{}   `json:"shipping_country"`
	ShippingPhoneNumber    interface{}   `json:"shipping_phone_number"`
	PaymentMethodType      string        `json:"payment_method_type"`
	Errors                 []interface{} `json:"errors"`
	Fingerprint            string        `json:"fingerprint"`
	VerificationValue      string        `json:"verification_value"`
	Number                 string        `json:"number"`
}

type PMJSON struct {
	PaymentMethod `json:"payment_method"`
}

type PMList struct {
	PMs []struct {
		Token             string        `json:"token"`
		CreatedAt         time.Time     `json:"created_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		GatewayType       string        `json:"gateway_type"`
		StorageState      string        `json:"storage_state"`
		ThirdPartyToken   string        `json:"third_party_token"`
		PaymentMethodType string        `json:"payment_method_type"`
		Errors            []interface{} `json:"errors"`
	} `json:"payment_methods"`
}

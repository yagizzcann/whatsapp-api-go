package go_whatsapp_official

type TextBasedTemplate struct {
	Name       string          `json:"name"`
	Language   Language        `json:"language"`
	Components []ComponentText `json:"components"`
	Type       string          `json:"-"`
	api        *API
}
type MultiBasedTemplate struct {
	Name       string           `json:"name"`
	Language   Language         `json:"language"`
	Components []ComponentMedia `json:"components"`
	Type       string           `json:"-"`
	api        *API
}
type Language struct {
	Code string `json:"code"`
}

type ComponentText struct {
	Type       string          `json:"type"`
	Parameters []ParameterText `json:"parameters"`
}

type ComponentMedia struct {
	Type       string           `json:"type"`
	Parameters []ParameterMedia `json:"parameters"`
}

type ParameterMedia struct {
	Type  string `json:"type"`
	Image struct {
		Link string `json:"link"`
	} `json:"image,omitempty"`
}

type ParameterText struct {
	Type     string   `json:"type"`
	Text     string   `json:"text,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	DateTime DateTime `json:"date_time,omitempty"`
}

type Currency struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    int    `json:"amount_1000"`
}

type DateTime struct {
	FallbackValue string `json:"fallback_value"`
}

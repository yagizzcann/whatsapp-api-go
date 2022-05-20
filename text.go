package whatsapp

type Text struct {
	Body       string `json:"body"`
	PreviewUrl bool   `json:"preview_url"`
	Type       string `json:"-"`
	api        *API
}

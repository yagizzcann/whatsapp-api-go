package whatsapp

type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Type      string `json:"-"`
	api       *API
}

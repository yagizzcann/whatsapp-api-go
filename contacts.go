package whatsapp

type ContactsReq struct {
	Contacts []Contacts `json:"contacts"`
	Type     string     `json:"-"`
	api      *API
}
type Contacts struct {
	Addresses []Address `json:"addresses"`
	Birthday  string    `json:"birthday"`
	Emails    []Email   `json:"emails"`
	Name      Name      `json:"name"`
	Org       Org       `json:"org"`
	Phones    []Phone   `json:"phones"`
	Urls      []Url     `json:"urls"`
}

type Address struct {
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Type        string `json:"type"`
}
type Email struct {
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Name struct {
	FormattedName string `json:"formatted_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MiddleName    string `json:"middle_name"`
	Suffix        string `json:"suffix"`
	Prefix        string `json:"prefix"`
}

type Org struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
}

type Phone struct {
	Phone string `json:"phone"`
	Type  string `json:"type"`
	WaID  string `json:"wa_id,omitempty"`
}

type Url struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}
